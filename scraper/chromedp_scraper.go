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

	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.OuterHTML("html", &res),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Simulate extraction logic
	// This is a simple example to show the extraction; you might need more complex parsing
	lines := strings.Split(res, "\n")
	for _, line := range lines {
		if strings.Contains(line, "<a href=") {
			// Extract the href and text content (very basic example)
			linkStart := strings.Index(line, "href=\"") + 6
			linkEnd := strings.Index(line[linkStart:], "\"") + linkStart

			// Check if indices are valid
			if linkStart < 6 || linkEnd < linkStart || linkEnd >= len(line) {
				continue
			}

			link := line[linkStart:linkEnd]

			textStart := strings.Index(line, ">") + 1
			textEnd := strings.Index(line[textStart:], "<") + textStart

			// Check if indices are valid
			if textStart < 1 || textEnd < textStart || textEnd >= len(line) {
				continue
			}

			text := line[textStart:textEnd]

			s.Data = append(s.Data, Data{Name: text, URL: link})
		}
	}

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
