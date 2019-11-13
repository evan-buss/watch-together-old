package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/evan-buss/watch-together/metadata/scraper/data"
// 	"github.com/evan-buss/watch-together/metadata/scraper/storage"
// )

// // Scraper defines the rules for a specific scraper
// type Scraper struct {
// 	Seed      []string
// 	Client    http.Client
// 	Wait      time.Duration  // The time to wait between requests
// 	Time      time.Duration  // The total time the scraper should run
// 	Writer    storage.Writer // A writer interface to output the results to storage
// 	UserAgent string
// 	Cancel    chan bool // Signals received from external factors
// 	Done      chan bool // Signals received from internal factors to notify outside listeners
// 	jobBuffer []string
// 	parser    data.ImdbData // The data model / must implement its own parsing logic
// }

// // Start takes a Scraper object and begins extracting data
// func (scraper *Scraper) Start() {

// 	results := make(chan data.Parser, 1) // Workers send back results from each job
// 	links := make(chan []string, 1)      // Workers send back the links they have found

// 	// Set the time that the crawler should run for. Send signal when limit hit
// 	var duration time.Duration
// 	if scraper.Time == -1 {
// 		duration = time.Hour * 100 // If user doesn't set limit, just use large value
// 	} else {
// 		duration = scraper.Time
// 	}

// 	// Set up the data storage model. Retrieve any urls that still need to be parsed
// 	cont, err := scraper.Writer.Init()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	scraper.Seed = append(scraper.Seed, cont...)

// 	stop := time.Tick(duration)
// 	links <- scraper.Seed // Send the seed links to the scheduler

// 	go scraper.scheduler(links)
// 	go scraper.dispatcher(results)

// 	for {
// 		select {
// 		case obj := <-results: // Worker returns a new Parser object
// 			// We now let the writer itself handle uniqueness checks and do as it sees fit
// 			if err := scraper.Writer.Write(obj); err != nil {
// 				fmt.Println(err)
// 			} else {
// 				links <- scraper.Writer.GetUnvisitedLinks(obj.GetLinks())
// 			}
// 		case <-stop:
// 			fmt.Println("Time limit hit")
// 			scraper.Writer.Close()
// 			scraper.Done <- true
// 			return
// 		case <-scraper.Cancel:
// 			fmt.Println("Signal cancel")
// 			scraper.Writer.Close()
// 			scraper.Done <- true
// 			return
// 		}
// 	}
// }

// // Scheduler simply reads all links from each parsed page and adds them to the work queue
// func (scraper *Scraper) scheduler(links <-chan []string) {
// 	for {
// 		select {
// 		case linkSlice := <-links:
// 			for _, url := range linkSlice {
// 				scraper.jobBuffer = append(scraper.jobBuffer, url)
// 			}
// 		}
// 	}
// }

// // dispatcher is responsible for keeping a feed of jobs sent to the extraction workers
// // It must regulate timing so that the server is not overloaded
// func (scraper *Scraper) dispatcher(results chan data.Parser) {

// 	numWaits := 0

// 	for {
// 		if len(scraper.jobBuffer) > 0 {
// 			url := scraper.jobBuffer[len(scraper.jobBuffer)-1]
// 			// Only scrape if we haven't visited the site before
// 			if !scraper.Writer.GetVisited(url) {
// 				time.Sleep(scraper.Wait)
// 				go scraper.extract(url, results)
// 			}
// 			scraper.jobBuffer = scraper.jobBuffer[:len(scraper.jobBuffer)-1]
// 			numWaits = 0
// 		} else {
// 			numWaits++
// 			time.Sleep(time.Millisecond * 100)
// 		}

// 		if numWaits > 20 {
// 			log.Println("Wait Timeout")
// 			scraper.Cancel <- true
// 		}
// 	}
// }

// // Scrape loads a specific url and scrapes data from it
// func (scraper *Scraper) extract(url string, results chan<- data.Parser) {
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		log.Println("Create Request Error")
// 		return
// 	}

// 	if scraper.UserAgent != "" {
// 		req.Header.Add("User-Agent", scraper.UserAgent)
// 	}

// 	resp, err := scraper.Client.Do(req)
// 	if err != nil {
// 		log.Println(url + " - " + err.Error())
// 		return
// 	}

// 	if resp.StatusCode != 200 {
// 		log.Println(url + " - " + strconv.Itoa(resp.StatusCode))
// 		return
// 	}

// 	if resp.StatusCode == 503 {
// 		log.Println("ERROR: 503. The server doesn't like our request traffic. Try to slow it down.")
// 		scraper.Cancel <- true
// 		return
// 	}

// 	body := resp.Body
// 	defer body.Close()

// 	obj, err := scraper.parser.Parse(body, url)
// 	if err == nil {
// 		results <- obj
// 	}
// }
