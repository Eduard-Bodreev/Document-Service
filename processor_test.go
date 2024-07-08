package main

import (
	"testing"
)

func TestDocumentProcessor(t *testing.T) {
	processor := NewDocumentProcessor()

	doc1 := &TDocument{Url: "http://example.com", PubDate: 100, FetchTime: 1, Text: "First"}
	doc2 := &TDocument{Url: "http://example.com", PubDate: 200, FetchTime: 2, Text: "Second"}
	doc3 := &TDocument{Url: "http://example.com", PubDate: 300, FetchTime: 3, Text: "Third"}

	processor.Process(doc1)
	processor.Process(doc2)
	result, _ := processor.Process(doc3)

	if result.Text != "Third" {
		t.Errorf("expected Text to be 'Third', got %s", result.Text)
	}
	if result.PubDate != 100 {
		t.Errorf("expected PubDate to be 100, got %d", result.PubDate)
	}
	if result.FirstFetchTime != 1 {
		t.Errorf("expected FirstFetchTime to be 1, got %d", result.FirstFetchTime)
	}
}
