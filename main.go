package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Environment Variables
//  - VIDEO_DIR

var videoDir = os.Getenv("VIDEO_DIR")

func main() {

	fmt.Println(videoDir)

	// Serve the site and static assets
	// http.Handle("/", http.FileServer(http.Dir("./site")))
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/media/", streamHandler)
	http.HandleFunc("/transcode/", transcodeHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexPage(w http.ResponseWriter, r *http.Request) {

	// transcodeCmd := exec.Command("bash", )

	http.ServeFile(w, r, "./site/index.html")
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	file := filepath.Join(videoDir, r.URL.Path[len("/media/"):])
	http.ServeFile(w, r, file)
}

func transcodeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("STARTING TRANSCODE")
	go StartConvert()
}
