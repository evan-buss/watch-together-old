package data

import (
	"errors"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ImdbData contains data from a single page
type ImdbData struct {
	URL     string   `json:"url,omitempty" db:"url"`
	Title   string   `json:"title,omitempty" db:"title"`
	Year    int      `json:"year,omitempty" db:"year"`
	Rating  string   `json:"rating,omitempty" db:"rating"`
	Summary string   `json:"summary,omitempty" db:"summary"`
	Poster  string   `json:"poster,omitempty" db:"poster"`
	Links   []string `json:"-" sql:"-" db:"-"`
}

// Parse extracts IMDB specific data
func (data ImdbData) Parse(body io.ReadCloser, url string) (Parser, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}
	imdb := ImdbData{}

	imdb.URL = url

	links := make(map[string]bool)

	// Find the title links
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists {
			link, err := validateURL(link)
			if err == nil {
				_, pres := links[link]
				if !pres {
					links[link] = true
				}
			}
		}
	})

	keys := make([]string, 0, len(links))
	for k := range links {
		keys = append(keys, k)
	}
	imdb.Links = keys

	// Find the review items
	poster, _ := doc.Find("div.poster > a > img").Attr("src")

	imdb.Poster = poster

	imdb.Summary = strings.TrimSpace(doc.Find("div.summary_text").Text())

	imdb.Title = strings.TrimSpace(doc.Find("div.title_wrapper > h1").Text())

	year := doc.Find("span#titleYear > a").Text()
	if year != "" {
		imdb.Year, err = strconv.Atoi(year)
	}

	imdb.Rating = doc.Find("div.ratingValue > strong > span").Text()

	log.Println(imdb.Title)
	return imdb, nil
}

// GetKey returns the model's key
func (data ImdbData) GetKey() string {
	return data.URL
}

// GetLinks returns the model's parsed links
func (data ImdbData) GetLinks() []string {
	return data.Links
}

// ====================================
// Website Specific Helper Methods
// ====================================

// IMDB movie parser. Looks for "/title/" links
func validateURL(url string) (string, error) {
	var re = regexp.MustCompile(`^(/title/tt[0-9]{7})`)
	if re.MatchString(url) {
		return "https://imdb.com" + re.FindString(url), nil
	}
	return "", errors.New("not a title string")
}
