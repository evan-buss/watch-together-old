package scraper

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/evan-buss/watch-together/info/data"
	"github.com/evan-buss/watch-together/info/output"
)

// Scraper defines the rules for a specific scraper
type Scraper struct {
	Seed      []string
	Client    http.Client
	Time      time.Duration // The total time the scraper should run
	Wait      time.Duration // The time to wait between requests
	DataType  data.Parser   // The data model / must implement its own parsing logic
	Writer    output.Writer // A writer interface to output the results to storage
	UserAgent string
	jobBuffer []string
	Cancel    chan bool // Signals received from external factors
	Done      chan bool // Signals received from internal factors to notify outside listeners
}

// Start takes a Scraper object and begins extracting data
func (scraper *Scraper) Start() {

	//parsed := make(map[string]data.Parser) // Map to hold the results of each parse
	results := make(chan data.Parser, 1) // Workers send back results from each job
	links := make(chan []string, 1)

	var duration time.Duration
	if scraper.Time == -1 {
		duration = time.Hour * 100 // If user doesn't set limit, just use large value
	} else {
		duration = scraper.Time
	}

	go scraper.scheduler(links)

	links <- scraper.Seed // Send the seed links to the dispatcher

	go scraper.dispatcher(results)

	stop := time.Tick(duration)

	err := scraper.Writer.Init()
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case obj := <-results: // Worker returns a new Parser object
			// We now let the writer itself handle uniqueness checks and do as it sees fit
			if err := scraper.Writer.Write(obj); err != nil {
				fmt.Println(err)
			} else {
				links <- scraper.Writer.GetUnvisitedLinks(obj.GetLinks())
			}
		case <-stop:
			fmt.Println("Time limit hit")
			//scraper.dumpData(parsed)
			scraper.Writer.Close()
			scraper.Done <- true
			return
		case <-scraper.Cancel:
			fmt.Println("Signal cancel")
			//scraper.dumpData(parsed)
			scraper.Writer.Close()
			scraper.Done <- true
			return
		}
	}
}

// Scheduler simply reads all links from each parsed page and adds them to the work queue
func (scraper *Scraper) scheduler(links <-chan []string) {
	for {
		select {
		case linkSlice := <-links:
			for _, url := range linkSlice {
				scraper.jobBuffer = append(scraper.jobBuffer, url)
			}
		}
	}
}

// dispatcher is responsible for keeping a feed of jobs sent to the extraction workers
// It must regulate timing so that the server is not overloaded
func (scraper *Scraper) dispatcher(results chan data.Parser) {
	// This is the only place we block. This is to ensure we don't get 503'ed on a server
	waits := 0
	for {
		time.Sleep(scraper.Wait)
		waits++
		if len(scraper.jobBuffer) > 0 {
			go scraper.extract(scraper.jobBuffer[len(scraper.jobBuffer)-1], results)
			scraper.jobBuffer = scraper.jobBuffer[:len(scraper.jobBuffer)-1]
			waits = 0 // Reset waits once we do something
		}
		// If we haven't done anything for a while, reset
		if waits > 10 {
			fmt.Println("jobs ran dry. Stopping parse")
			scraper.Cancel <- true
		}
	}
}

// Scrape loads a specific url and scrapes data from it
func (scraper *Scraper) extract(url string, results chan<- data.Parser) {
	log.Println("Parsing", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Create Request Error")
		return
	}

	if scraper.UserAgent != "" {
		req.Header.Add("User-Agent", scraper.UserAgent)
	}

	resp, err := scraper.Client.Do(req)
	if err != nil {
		log.Println(url + " - " + err.Error())
		return
	}

	if resp.StatusCode != 200 {
		log.Println(url + " - " + strconv.Itoa(resp.StatusCode))
		return
	}

	if resp.StatusCode == 503 {
		log.Println("triggering shutdown")
		scraper.Cancel <- true
		return
	}

	body := resp.Body
	defer body.Close()

	obj, err := scraper.DataType.Parse(body, url)
	if err == nil {
		results <- obj
	}
}
