package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/evan-buss/watch-together/info/data"
	"github.com/evan-buss/watch-together/info/output"
	"github.com/evan-buss/watch-together/info/scraper"
)

var mode *string

//var format *string

func init() {
	mode = flag.String("mode", "quotes", "Which scraper to use (imdb or quotes)")
	//format = flag.String("o", "db", "Where to output results (json or db)")
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
		fmt.Println("quotes found")
		dataType = data.QuoteData{}
	} else if *mode == "imdb" {
		fmt.Println("imdb found")
		dataType = data.ImdbData{}
	}

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	crawler := scraper.Scraper{
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
			Insert: `INSERT OR IGNORE INTO movies (url, title, year, rating, summary, poster) 
			VALUES (?, ?, ?, ?, ?, ?)`,
		},
		Wait:      time.Millisecond * 400,
		Time:      time.Minute * 30,
		UserAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:70.0) Gecko/20100101 Firefox/70.0",
		DataType:  dataType,
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
