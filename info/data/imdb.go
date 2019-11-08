package data

import (
	"errors"
	"io"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ImdbData contains data from a single page
type ImdbData struct {
	URL     string   `json:"url"`
	Title   string   `json:"title"`
	Year    string   `json:"year"`
	Rating  string   `json:"rating"`
	Summary string   `json:"summary"`
	Poster  string   `json:"poster"`
	Links   []string `json:"-" sql:"-"`
}

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
	// log.Println(imdb.Title)

	imdb.Year = doc.Find("span#titleYear > a").Text()

	imdb.Rating = doc.Find("div.ratingValue > strong > span").Text()

	return imdb, nil
}

func (data ImdbData) GetKey() string {
	return data.URL
}

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
