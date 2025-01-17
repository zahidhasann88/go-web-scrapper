package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
)

type ScraperInterface interface {
	Scrape(url string)
	ExportData(format, filename string) error
}

type Scraper struct {
	Collector *colly.Collector
	Data      []Data
}

func NewScraper() *Scraper {
	c := colly.NewCollector()
	return &Scraper{Collector: c}
}

func (s *Scraper) Scrape(url string) {
	s.Data = []Data{}

	// OnHTML callback to scrape data
	s.Collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		s.Data = append(s.Data, Data{Name: e.Text, URL: link})
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	})

	// OnHTML callback to handle pagination
	s.Collector.OnHTML("a.next", func(e *colly.HTMLElement) {
		nextPage := e.Attr("href")
		e.Request.Visit(nextPage)
	})

	s.Collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	s.Collector.Visit(url)
	s.Collector.Wait()
}

func (s *Scraper) ExportData(format, filename string) error {
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
