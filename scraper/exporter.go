package scraper

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"os"
)

type Data struct {
	Name string `json:"name" xml:"name"`
	URL  string `json:"url" xml:"url"`
}

func ExportToCSV(data []Data, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Name", "URL"}) // write header
	for _, d := range data {
		writer.Write([]string{d.Name, d.URL})
	}
	return nil
}

func ExportToJSON(data []Data, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(data)
}

func ExportToXML(data []Data, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	return encoder.Encode(data)
}
