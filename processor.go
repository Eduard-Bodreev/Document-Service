package main

import (
	"sync"
)

type DocumentProcessor struct {
	mu      sync.Mutex
	docs    map[string]*TDocument
	pubDate map[string]uint64
	firstFT map[string]uint64
}

func NewDocumentProcessor() *DocumentProcessor {
	return &DocumentProcessor{
		docs:    make(map[string]*TDocument),
		pubDate: make(map[string]uint64),
		firstFT: make(map[string]uint64),
	}
}

func (dp *DocumentProcessor) Process(d *TDocument) (*TDocument, error) {
	dp.mu.Lock()
	defer dp.mu.Unlock()

	existingDoc, exists := dp.docs[d.Url]
	if !exists {
		dp.docs[d.Url] = d
		dp.pubDate[d.Url] = d.PubDate
		dp.firstFT[d.Url] = d.FetchTime
		return d, nil
	}

	if d.FetchTime > existingDoc.FetchTime {
		dp.docs[d.Url].Text = d.Text
		dp.docs[d.Url].FetchTime = d.FetchTime
	}

	if d.FetchTime < existingDoc.FetchTime {
		if dp.pubDate[d.Url] > d.PubDate {
			dp.pubDate[d.Url] = d.PubDate
		}
		if dp.firstFT[d.Url] > d.FetchTime {
			dp.firstFT[d.Url] = d.FetchTime
		}
	}

	dp.docs[d.Url].PubDate = dp.pubDate[d.Url]
	dp.docs[d.Url].FirstFetchTime = dp.firstFT[d.Url]

	return dp.docs[d.Url], nil
}
