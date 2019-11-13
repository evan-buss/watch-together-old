package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/evan-buss/watch-together/metadata/scraper/storage"
)

func main() {

	// Get seed urls from command arguments
	if len(os.Args) < 1 {
		fmt.Println("Usage: imdb_scraper [url]")
		os.Exit(1)
	}

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	crawler := Scraper{
		Seed: os.Args[1:],
		Client: http.Client{
			Timeout: time.Second * 2,
		},
		Writer:    &storage.SQLite{},
		Workers:   2,
		Time:      -1,
		UserAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:70.0) Gecko/20100101 Firefox/70.0",
		Cancel:    make(chan bool, 1),
		Done:      make(chan bool, 1),
	}

	// Wait for Ctrl+C signal to shutdown
	go func() {
		<-sigs
		fmt.Println("Cancel signal received.")
		crawler.Cancel <- true
	}()

	go crawler.Start()

	select {
	case <-crawler.Done:
		fmt.Println("Exit signal recieved...")
	}
}
