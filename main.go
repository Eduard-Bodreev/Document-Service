package main

import (
	"fmt"
)

type TDocument struct {
	Url            string
	PubDate        uint64
	FetchTime      uint64
	Text           string
	FirstFetchTime uint64
}

func main() {
	processor := NewDocumentProcessor()

	doc1 := &TDocument{Url: "http://example.com", PubDate: 100, FetchTime: 1, Text: "First"}
	doc2 := &TDocument{Url: "http://example.com", PubDate: 200, FetchTime: 2, Text: "Second"}
	doc3 := &TDocument{Url: "http://example.com", PubDate: 300, FetchTime: 3, Text: "Third"}

	processor.Process(doc1)
	processor.Process(doc2)
	result, _ := processor.Process(doc3)

	fmt.Printf("Processed Document: %+v\n", result)
}
