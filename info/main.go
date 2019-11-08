package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/evan-buss/scaper/data"
	"github.com/evan-buss/scaper/output"
	"github.com/evan-buss/scaper/scraper"
)

var mode *string

func init() {
	mode = flag.String("mode", "quotes", "Which scraper to use (imdb or quotes)")
}

func main() {
	flag.Parse()
	// Get starting url
	if len(flag.Args()) < 1 {
		fmt.Println("Usage: imdb_scraper [url]")
		os.Exit(1)
	}

	var dataType data.Parser
	if *mode == "quotes" {
		dataType = data.QuoteData{}
	} else if *mode == "imdb" {
		dataType = data.ImdbData{}
	}

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	scraper := scraper.Scraper{
		Seed: flag.Args(),
		Client: http.Client{
			Timeout: time.Second * 10,
		},
		Writer: &output.SQLite{
			DataType: dataType,
			Create: `CREATE TABLE IF NOT EXISTS movies (
				url TEXT PRIMARY KEY,
				title TEXT,
				year TEXT,
				rating TEXT,
				summary TEXT,
				poster TEXT
			)`,
			Insert: `INSERT INTO movies 
			(url, title, year, rating, summary, poster) VALUES (?, ?, ?, ?, ?, ?)`,
		},
		Wait:     time.Millisecond * 300,
		Time:     time.Minute * 30,
		DataType: dataType,
		Cancel:   make(chan bool, 1),
		Done:     make(chan bool, 1),
	}

	go func() {
		<-sigs
		fmt.Println("Cancel signal received.")
		scraper.Cancel <- true
	}()

	go scraper.Start()

	select {
	case <-scraper.Done:
		fmt.Println("Exit signal recieved...")
	}
}
