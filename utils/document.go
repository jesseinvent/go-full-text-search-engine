package utils

import (
	"compress/gzip"
	"encoding/xml"
	"log"
	"os"
)

type Document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

// Loads documents from path
func LoadDocuments(path string) ([]Document, error) {

	// Load file
	f, err := os.Open(path)

	if err != nil {
		log.Printf("Error: could not load file on path %s", path)
		return nil, err
	}

	defer f.Close()

	// Read zipped file
	gz, err := gzip.NewReader(f)

	if err != nil {
		log.Printf("Error: could not extract Gzip file on path %s", path)
		return nil, err
	}

	defer gz.Close()

	// Decode xml file
	doc := xml.NewDecoder(gz)

	dump := struct {
		Documents []Document `xml:"doc"`
	}{}

	if err := doc.Decode(&dump); err != nil {
		log.Printf("Error: could not decode file")
		return nil, err
	}

	docs := dump.Documents

	// Loop through documents and assign ID
	for i := range docs {
		docs[i].ID = i
	}

	return docs, nil
}
