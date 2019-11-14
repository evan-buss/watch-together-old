package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
)

// Metadata contains the results from running FFProbe on a single file
type Metadata struct {
	Streams []Stream `json:"streams"`
	Format  Format   `json:"format"`
}

// Stream contains information about a single stream
type Stream struct {
	Index     int    `json:"index"`
	CodecName string `json:"codec_name"`
	Profile   string `json:"profile"`
	CodecType string `json:"codec_type"`
}

// Format contains information about a single file
type Format struct {
	FileName       string     `json:"filename"`
	NbStreams      int        `json:"nb_streams"`
	FormatName     string     `json:"format_name"`
	FormatLongName string     `json:"format_long_name"`
	Duration       string     `json:"duration"`
	Size           string     `json:"size"`
	Bitrate        string     `json:"bit_rate"`
	ProbeScore     int        `json:"probe_scrore"`
	Tags           FormatTags `json:"tags"`
}

// FormatTags gives extra info about the file format
type FormatTags struct {
	Title string `json:"title"`
}

// MovieDBInfo represents a single movie returned from the Movie Metadata API
type MovieDBInfo struct {
	Poster  string `json:"poster"`
	Rating  string `json:"rating"`
	Summary string `json:"summary"`
	Title   string `json:"title"`
	Year    string `json:"year"`
}

var titleMatcher = regexp.MustCompile(`(^.+)\s\((.+)\)`)

// FFProbe runs the FFProbe utility on the given file. It attempts to load the metadata from the metadata REST API
//  based on the filename
// Movie filenames should follow the
func FFProbe(path string, info os.FileInfo) {
	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", path)

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	// Parse the FFProbe command into Metadata object
	var metadata Metadata
	err = json.Unmarshal(out, &metadata)
	if err != nil {
		log.Println(err)
	}
	var queries []string
	if titleMatcher.Match([]byte(info.Name())) {
		queries = titleMatcher.FindStringSubmatch(info.Name())
	} else if titleMatcher.Match([]byte(metadata.Format.Tags.Title)) {
		queries = titleMatcher.FindStringSubmatch(metadata.Format.Tags.Title)
	} else {
		return
	}

	baseURL, err := url.Parse("http://localhost:8080/")
	if err != nil {
		log.Println(err)
		return
	}

	params := url.Values{}
	params.Add("title", queries[1])
	params.Add("year", queries[2])
	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		log.Println("get", err)
	}

	if resp.StatusCode != 200 {
		fmt.Println("non 200")
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var movieInfo []map[string]interface{}
	err = json.Unmarshal(data, &movieInfo)
	if err != nil {
		log.Println("json", err)
	}

	fmt.Printf("%+v\n", movieInfo)
}
