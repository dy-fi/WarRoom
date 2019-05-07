package scrapper

import (
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"

	"github.com/dy-fi/war-room/models"

	"strings"
	"time"
)

// ScrapeAgent builds a scrapper agent and gets all location values from that URL model
func ScrapeAgent(url models.URL) (map[string]string, error) {
	// init collector
	var s = colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.Async(true),
		colly.AllowedDomains("url"),
	)

	// limit rules
	s.Limit(&colly.LimitRule{
		// tweak later if needed
		Delay:       1 * time.Second,
		RandomDelay: 1 * time.Second,
	})

	// begin collection
	s.Visit(url.Glob)

	var b string
	// assign b to the body string 
	s.OnResponse(func(r *colly.Response) {
		b = string(r.Body)
	})

	// result map
	m := make(map[string]string)
	// scrape for every location in URL
	for _, i := range url.Locations {
		value, err := findElement(b, i.Address)
		if err != nil {
			return m, err
		}
		m[i.Key] = value
	}

	return m, nil
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
