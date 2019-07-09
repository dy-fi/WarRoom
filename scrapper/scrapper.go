package scrapper

import (
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"

	"github.com/dy-fi/war-room/backend/models"

	"strings"
	"time"
)

// ScrapeData is scraped data from the webpage
type ScrapeData struct {
	name	string	
	data	string
}

// ScrapeAgent builds a scrapper agent and gets all location values from that URL model
func ScrapeAgent(page models.Page) ([]ScrapeData) {
	// init collector
	var s = colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.Async(true),
		colly.AllowedDomains(page.Glob),
	)

	// limit rules
	s.Limit(&colly.LimitRule{
		// tweak later if needed
		Delay:       1 * time.Second,
		RandomDelay: 1 * time.Second,
	})

	// begin collection
	s.Visit(page.Glob)

	var b string
	// assign b to the body string 
	s.OnResponse(func(r *colly.Response) {
		b = string(r.Body)
	})

	// result list
	m := []ScrapeData{}
	// scrape for every location in URL
	for _, i := range page.Locations {
		value, err := findElement(b, i.Address)
		if err != nil {
			m = append(m, ScrapeData{i.Key, "ERROR"})
		}
		m = append(m, ScrapeData{i.Key, value})
	}

	return m
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
