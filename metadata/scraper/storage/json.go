package storage

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/evan-buss/watch-together/metadata/scraper/data"
)

// JSON is an output.Writer that saves data to a JSON file
type JSON struct {
	DataType data.Parser
	file     *os.File
	store    map[string]bool
}

// Init creates the JSON file and gets it ready
func (j *JSON) Init() error {
	// We get the actual data type of the given data.Parser to use as json file name
	name := reflect.TypeOf(j.DataType).String()
	name = name[strings.Index(name, ".")+1:]
	file, err := os.Create(name + ".json")
	if err != nil {
		return err
	}
	j.file = file

	//Initialize data store to prevent duplicates
	j.store = make(map[string]bool)

	_, err = file.WriteString("[")
	if err != nil {
		return err
	}
	return nil
}

// Write writes a single data.Parser object to the JSON file
func (j *JSON) Write(obj data.Parser) error {

	_, pres := j.store[obj.GetKey()]
	// Object is already present. Skip
	if pres {
		return errors.New("object already exists")
	}

	objJSON, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	_, err = j.file.Write(objJSON)
	if err != nil {
		return err
	}
	_, err = j.file.WriteString(",")
	if err != nil {
		return err
	}

	// Only store it in memory it once it is successfully saved to storage file
	j.store[obj.GetKey()] = true

	return nil
}

// GetUnvisitedLinks checks the in memory store for each url and returns an array of urls that haven't been visited before
func (j *JSON) GetUnvisitedLinks(links []string) []string {
	out := make([]string, 0)
	for _, url := range links {
		_, pres := j.store[url]
		if !pres {
			out = append(out, url)
		}
	}
	return out
}

// GetVisited returns true if the provided url has already been visited
func (j *JSON) GetVisited(url string) bool {
	_, pres := j.store[url]
	return pres
}

// Close finalizes the JSON files and closes it
func (j *JSON) Close() {
	var lastChar [1]byte
	_, err := j.file.Read(lastChar[:])
	if err != nil {
		log.Println(err)
	}
	if string(lastChar[:]) != "]" {
		_, err := j.file.Seek(-1, 1) //Remove extra comma at the end
		if err != nil {
			log.Println(err)
		}
		_, err = j.file.WriteString("]")
		if err != nil {
			log.Println(err)
		}
	}
	err = j.file.Sync()
	if err != nil {
		log.Println(err)
	}
	err = j.file.Close()
	if err != nil {
		log.Println(err)
	}
}
