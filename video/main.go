package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var videoPath = "/home/evan/Videos"

func main() {
	err := filepath.Walk(videoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		if info.IsDir() {
			// fmt.Println("DIR:", info.Name())
		} else {
			// fmt.Println("FILE:", info.Name())
			if strings.Contains(info.Name(), ".mkv") {
				FFProbe(path, info)
			}
		}
		return nil
	})
	log.Println(err)
}
