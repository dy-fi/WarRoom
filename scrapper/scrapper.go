package scrapper

import (
	"github.com/antchfx/htmlquery"

	"fmt"
	// "os"
	// "strings"
	// "time"
)

// ScrapedData is scraped data from the webpage
type ScrapedData struct {
	// Name reference
	Name string `json:"name"`
	// Data body
	Data string `json:"data"`
}

func scrapeXPath(url, path string) (string, error) {
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		return "-1", err
	}

	element := htmlquery.FindOne(doc, path)
	text := ""

	if element != nil {
		text = htmlquery.InnerText(element)
	}

	return text, nil
}

func main() {
	url := "https://weather.com/weather/today/l/c1535f42ba5fc52449e416514aca69b3b2a16aae4b89abd6c92e662f7a89c02f"
	xpath := "//*[@id=\"main-Nowcard-92c6937d-b8c3-4240-b06c-9da9a8b0d22b\"]/div/div/section/div[3]/table/tbody/tr[1]/td/span"
	data, err := scrapeXPath(url, xpath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response: \n")
	fmt.Print(data)
}
