package repos

import (
	// "github.com/robfig/cron"

	"github.com/dy-fi/war-room/models"
	"github.com/dy-fi/war-room/scrapper"
)

// ScrapeCity for all needed information
func ScrapeCity(c *models.City) []scrapper.ScrapedData {
	cityData := []scrapper.ScrapedData{}
	for _, i := range c.Places {
		s, err := scrapper.Agent(i.URL, i.Address)
		if err != nil {
			scrapeError := scrapper.ScrapedData{i.Name, err.Error()}
			cityData = append(cityData, scrapeError)
		}
		data := scrapper.ScrapedData{i.Name, s}
		cityData = append(cityData, data)
	}

	return cityData
}

// func cityCron(c *models.City, interval int) []scrapper.ScrapeData {
// 	// cron init
// 	cron := cron.New()
// 	defer cron.Stop()

// 	c.AddFunc()
// }
