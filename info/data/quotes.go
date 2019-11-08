package data

import (
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
)

// Practice scraper for quotes.toscrape.com

// QuoteData holds information from a single quote page
type QuoteData struct {
	URL    string   `json:"url"`
	Quotes []Quote  `json:"quotes"`
	Links  []string `json:"links"`
}

// Quote holds information about a single quote
type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

func (data QuoteData) Parse(body io.ReadCloser, url string) (Parser, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}
	quote := QuoteData{}
	quote.URL = url

	quote.Quotes = make([]Quote, 0)

	doc.Find("div.quote").Each(func(i int, s *goquery.Selection) {
		text := s.Find("span.text").Text()
		author := s.Find("span > small.author").Text()
		quote.Quotes = append(quote.Quotes, Quote{text, author})
	})

	link, exists := doc.Find("li.next > a").Attr("href")
	if exists {
		quote.Links = append(quote.Links, "http://quotes.toscrape.com"+link)
	}

	return quote, nil
}

func (data QuoteData) GetKey() string {
	return data.URL
}

func (data QuoteData) GetLinks() []string {
	return data.Links
}
