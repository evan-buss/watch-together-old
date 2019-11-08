package output

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"

	"github.com/evan-buss/watch-together/info/scraper/data"
)

// JSON is an output.Writer that saves data to a JSON file
type JSON struct {
	DataType data.Parser
	File     *os.File
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
	j.File = file

	file.WriteString("[")
	return nil
}

// WriteSingle writes a single data.Parser object to the JSON file
func (j *JSON) WriteSingle(obj data.Parser) error {
	json, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	j.File.Write(json)
	j.File.WriteString(",")
	return nil
}

// WriteFull writes an array of data.Parser objects to the JSON file
func (j *JSON) WriteFull(objs []data.Parser) error {
	json, err := json.Marshal(objs)
	if err != nil {
		return err
	}

	j.File.Write(json)
	return nil
}

// Close finalizes the JSON files and closes it
func (j *JSON) Close() {
	var lastChar [1]byte
	j.File.Read(lastChar[:])
	if string(lastChar[:]) != "]" {
		j.File.Seek(-1, 1) //Remove extra comma at the end
		j.File.WriteString("]")
	}
	j.File.Sync()
	j.File.Close()
}
