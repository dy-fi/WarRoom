package scrapper

import (
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"

	"strings"
	"time"
)

// ScrapeData is scraped data from the webpage
type ScrapeData struct {
	Name string
	Data string
}

// ScrapeAgent builds a scrapper agent and gets all location values from that URL model
func ScrapeAgent(url string, address string) (string, error) {
	// init collector
	var s = colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.Async(true),
		colly.AllowedDomains(url),
	)

	// limit rules
	s.Limit(&colly.LimitRule{
		// tweak later if needed
		Delay:       1 * time.Second,
		RandomDelay: 1 * time.Second,
	})

	// begin collection
	s.Visit(url)

	var b string
	// assign b to the body string
	s.OnResponse(func(r *colly.Response) {
		b = string(r.Body)
	})

	value, err := findElement(b, address)
	if err != nil {
		return err.Error(), err
	}
	return value, nil
}

func findElement(body, xpath string) (string, error) {
	doc, err := htmlquery.Parse(strings.NewReader(body))
	if err != nil {
		return "-1", err
	}

	element := htmlquery.FindOne(doc, xpath)
	text := htmlquery.InnerText(element)

	return text, nil
}
