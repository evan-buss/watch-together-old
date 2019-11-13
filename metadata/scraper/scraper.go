package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/evan-buss/watch-together/metadata/scraper/data"
	"github.com/evan-buss/watch-together/metadata/scraper/storage"
)

// Scraper defines the rules for a specific scraper
type Scraper struct {
	Seed   []string
	Client http.Client
	// The number of concurrent requests to make. Workers
	// wait for all others to finish before processing next batch of work
	Workers   int
	Time      time.Duration  // The total time the scraper should run
	Writer    storage.Writer // A writer interface to output the results to storage
	UserAgent string
	Cancel    chan bool // Signals received from external factors
	Done      chan bool // Signals received from internal factors to notify outside listeners
	jobBuffer []string
	parser    data.ImdbData // The data model / must implement its own parsing logic
	wg        sync.WaitGroup
}

// Start takes a Scraper object and begins extracting data
func (scraper *Scraper) Start() {

	results := make(chan data.Parser) // Workers send back results from each job
	jobs := make(chan string)
	// links := make(chan []string, 1)      // Workers send back the links they have found

	// Set the time that the crawler should run for. Send signal when limit hit
	var duration time.Duration
	if scraper.Time == -1 {
		duration = time.Hour * 100 // If user doesn't set limit, just use large value
	} else {
		duration = scraper.Time
	}

	// Set up the data storage model. Retrieve any urls that still need to be parsed
	cont, err := scraper.Writer.Init()
	if err != nil {
		log.Fatal(err)
	}
	scraper.Seed = append(scraper.Seed, cont...)
	scraper.jobBuffer = append(scraper.jobBuffer, scraper.Seed...)

	// Launch the workers
	for i := 0; i < scraper.Workers; i++ {
		go scraper.worker(i, jobs, results)
	}
	go scraper.buffer(jobs) // We need this in a couroutine because otherwise we don't recieve any results until all seeds are processed. Very long blocking
	go scraper.receiver(jobs, results)

	stop := time.Tick(duration)
	for {
		select {
		case <-stop:
			scraper.Writer.Close()
			scraper.Done <- true
		case <-scraper.Cancel:
			scraper.Writer.Close()
			scraper.Done <- true
		}
	}
}

// buffer provides a constant stream of jobs sends the initial jobs to the queue
func (scraper *Scraper) buffer(jobs chan<- string) {
	lastActive := time.Now()
	// We have a five second wait period before we shut down the entire parser.
	for time.Now().Sub(lastActive) < (time.Second * 5) {
		if len(scraper.jobBuffer) > 0 {
			url := scraper.jobBuffer[len(scraper.jobBuffer)-1]
			// Check again before processing that we haven't already visited it
			// If there are duplicates in the queue, we won't know that until after we
			// process the first one
			if !scraper.Writer.GetVisited(url) {
				jobs <- url
			}
			scraper.jobBuffer = scraper.jobBuffer[:len(scraper.jobBuffer)-1]
			// Once we have actuall processed a job, reset the wait timer
			lastActive = time.Now()
		}
	}
	// The buffer has been idle for over the wait time, shut the crawler down
	close(jobs)
	scraper.Cancel <- true
}

// Receiver sends results to the writer and appends links to the job buffer
func (scraper *Scraper) receiver(jobs chan<- string, results <-chan data.Parser) {
	for {
		obj := <-results
		if err := scraper.Writer.Write(obj); err != nil {
			fmt.Println(err)
		} else {
			links := scraper.Writer.GetUnvisitedLinks(obj.GetLinks())
			scraper.jobBuffer = append(scraper.jobBuffer, links...)
		}
	}
}

// worker listens for jobs and extracts data from the url
func (scraper *Scraper) worker(id int, jobs <-chan string, results chan<- data.Parser) {
	for url := range jobs {
		scraper.wg.Add(1)
		obj, err := scraper.extract(url)
		if err != nil {
			log.Println(err)
			continue
		}
		results <- obj
		// Wait for all extractions to finish before we go again
		scraper.wg.Wait()
	}
}

// Scrape loads a specific url and scrapes data from it
func (scraper *Scraper) extract(url string) (data.Parser, error) {
	log.Println("EXTRACT:", url)
	scraper.wg.Done()

	// reqStart := time.Now()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Create Request Error")
		return nil, err
	}

	if scraper.UserAgent != "" {
		req.Header.Add("User-Agent", scraper.UserAgent)
	}

	resp, err := scraper.Client.Do(req)
	if err != nil {
		log.Println(url + " - " + err.Error())
		return nil, err
	}
	// fmt.Println(time.Now().Sub(reqStart))

	if resp.StatusCode == 503 {
		scraper.Cancel <- true
		return nil, errors.New("503 response received. slow down boss")
	}

	if resp.StatusCode != 200 {
		log.Println(url + " - " + strconv.Itoa(resp.StatusCode))
		return nil, errors.New("non 200 response code recieved")
	}

	body := resp.Body
	defer body.Close()

	obj, err := scraper.parser.Parse(body, url)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
