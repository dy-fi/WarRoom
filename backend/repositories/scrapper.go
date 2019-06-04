package repos

import (
	// "github.com/robfig/cron"

	"github.com/dy-fi/war-room/models"
	"github.com/dy-fi/war-room/scrapper"
)


// ScrapeCity for all needed information
func ScrapeCity(c *models.City) []scrapper.ScrapeData {
	cityData := []scrapper.ScrapeData{}
	for _, i := range c.Pages {
		s := scrapper.ScrapeAgent(i)
		cityData = append(cityData, s...)
	}

	return cityData
}

// func cityCron(c *models.City, interval int) []scrapper.ScrapeData {
// 	// cron init
// 	cron := cron.New()
// 	defer cron.Stop()
	
// 	c.AddFunc()
// }