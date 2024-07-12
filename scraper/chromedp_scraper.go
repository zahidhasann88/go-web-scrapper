package scraper

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
)

type ChromedpScraper struct {
	Data []Data
}

func NewChromedpScraper() *ChromedpScraper {
	return &ChromedpScraper{}
}

func (s *ChromedpScraper) Scrape(url string) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Function to extract links and follow pagination
	var scrapePage func(string)
	scrapePage = func(pageURL string) {
		var res string
		err := chromedp.Run(ctx,
			chromedp.Navigate(pageURL),
			chromedp.OuterHTML("html", &res),
		)
		if err != nil {
			log.Fatal(err)
		}

		// Extract links
		lines := strings.Split(res, "\n")
		for _, line := range lines {
			if strings.Contains(line, "<a href=") {
				linkStart := strings.Index(line, "href=\"") + 6
				linkEnd := strings.Index(line[linkStart:], "\"") + linkStart
				if linkStart < 6 || linkEnd < linkStart || linkEnd >= len(line) {
					continue
				}
				link := line[linkStart:linkEnd]

				textStart := strings.Index(line, ">") + 1
				textEnd := strings.Index(line[textStart:], "<") + textStart
				if textStart < 1 || textEnd < textStart || textEnd >= len(line) {
					continue
				}
				text := line[textStart:textEnd]

				s.Data = append(s.Data, Data{Name: text, URL: link})
			}
		}

		// Follow pagination
		for _, line := range lines {
			if strings.Contains(line, "class=\"next\"") {
				nextPageStart := strings.Index(line, "href=\"") + 6
				nextPageEnd := strings.Index(line[nextPageStart:], "\"") + nextPageStart
				if nextPageStart < 6 || nextPageEnd < nextPageStart || nextPageEnd >= len(line) {
					continue
				}
				nextPage := line[nextPageStart:nextPageEnd]
				log.Printf("Following next page: %s\n", nextPage)
				scrapePage(nextPage)
				break
			}
		}
	}
	log.Printf("Starting scraping for URL: %s\n", url)
	scrapePage(url)
	log.Printf("Scraping completed. Extracted %d data entries.\n", len(s.Data))
	fmt.Println(s.Data)
}

func (s *ChromedpScraper) ExportData(format, filename string) error {
	switch format {
	case "csv":
		return ExportToCSV(s.Data, filename)
	case "json":
		return ExportToJSON(s.Data, filename)
	case "xml":
		return ExportToXML(s.Data, filename)
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
}
