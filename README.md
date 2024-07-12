# Go Web Scraper

This project demonstrates a web scraping API using Golang with Gin and ChromeDP for dynamic site scraping.

## Setup

### Prerequisites

- Go (version 1.16+ recommended)
- Git
- Chrome or Chromium browser (for ChromeDP scraper)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/zahidhasann88/go-web-scraper.git
   cd go-web-scraper
2. Install dependencies:
   ```bash
   go mod tidy

3. Run the application
   ```bash
   go run main.go
   
## Usage

### API Endpoint

- **Endpoint:** `POST /scrape`
- **Description:** Scrapes a website using ChromeDP or Colly based on the `useChromedp` flag.

### Example Request
- **POST** - http://localhost:8080/scrape
```json

{
    "url": "https://executivemachines.com",
    "format": "json",
    "filename": "scraped_data.json",
    "useChromedp": true
}
```
# Technologies Used

This project utilizes the following technologies:

- **Gin** - Web framework for building APIs in Golang.
- **ChromeDP** - Headless Chrome DevTools Protocol for browser automation and scraping.
- **Colly** - Golang-based web scraping framework for extracting data from websites.
