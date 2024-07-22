package main

import (
	"flag"
	"log"
	"time"

	"github.com/jesseinvent/go-full-text-search/utils"
)

func main() {

	var dumpPath, query string

	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	flag.StringVar(&query, "q", "Small wild cat", "search query")
	flag.Parse()

	print("Full text search is in progress...")

	start := time.Now()

	// Load Document
	docs, err := utils.LoadDocuments(dumpPath)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	// Create Index
	start = time.Now()
	index := make(utils.Index)

	// Add documents to index
	index.Add(docs)

	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	// Search though index for matched documents
	start = time.Now()
	matchedIDs := index.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))

	// Loop through all matched IDs
	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}

}
