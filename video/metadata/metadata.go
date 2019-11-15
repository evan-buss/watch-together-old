package metadata

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

var titleMatcher = regexp.MustCompile(`(^.+)\s\((.+)\)`)

// ParseDir walks the given directory in search of video files. Each video file encountered is saved to the
// metadata database
func ParseDir(dir string) {
	fmt.Println("Scanning", viper.GetString("video-dir"), "for movie files.")
	createStore()
	total := 0
	matched := 0

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		if !info.IsDir() {
			// fmt.Println("DIR:", info.Name())
			for _, ext := range movieFormats {
				// Loop through all known extensions until we get a match
				if strings.Contains(info.Name(), ext) {
					total++
					err := FFProbe(path, info)
					if err == nil {
						matched++
					}
					break
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Parsed Found", total, "movies")
	fmt.Println("Found metadata matches for", matched, "movies")
	fmt.Println(total-matched, "movies must be manually assigned metadata")
	db.Close()
}

// FFProbe runs the FFProbe utility on the given file. It attempts to load the metadata from the metadata REST API
//  based on the filename
// Movie filenames should follow the
func FFProbe(path string, info os.FileInfo) error {
	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", path)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	// Parse the FFProbe command into Metadata object
	var metadata Metadata
	err = json.Unmarshal(out, &metadata)
	if err != nil {
		return err
	}
	var queries []string
	if titleMatcher.Match([]byte(info.Name())) {
		queries = titleMatcher.FindStringSubmatch(info.Name())
	} else if titleMatcher.Match([]byte(metadata.Format.Tags.Title)) {
		queries = titleMatcher.FindStringSubmatch(metadata.Format.Tags.Title)
	} else {
		return errors.New("Couldn't parse")
	}

	baseURL, err := url.Parse("http://localhost:8080/")
	if err != nil {
		return err
	}

	params := url.Values{}
	params.Add("title", queries[1])
	params.Add("year", queries[2])
	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		fmt.Println("non 200")
		return errors.New("Non 200 Status Code")
	}

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var movieInfo []MovieDBInfo
	err = json.Unmarshal(data, &movieInfo)
	if err != nil {
		return err
	}
	if len(movieInfo) > 0 {
		_ = db.MustExec(`INSERT INTO movies (location, metadata) VALUES (?, ?)`, movieInfo[0].Title, movieInfo[0].RowID)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("Movie not found in database")
}
