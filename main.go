package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zahidhasann88/go-web-scraper/scraper"
)

type ScrapeRequest struct {
	URL         string `json:"url" binding:"required"`
	Format      string `json:"format" binding:"required"`
	Filename    string `json:"filename" binding:"required"`
	UseChromedp bool   `json:"useChromedp"` // Add a flag for using chromedp
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the Golang Web Scraper!")
	})

	r.POST("/scrape", func(c *gin.Context) {
		var req ScrapeRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var s scraper.ScraperInterface
		if req.UseChromedp {
			s = scraper.NewChromedpScraper()
		} else {
			s = scraper.NewScraper()
		}

		s.Scrape(req.URL)

		if err := s.ExportData(req.Format, req.Filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error exporting data: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Scraping and export completed!"})
	})

	r.Run(":8080")
}
