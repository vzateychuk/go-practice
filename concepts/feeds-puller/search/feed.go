package search

import (
	"encoding/json"
	"os"
)

// Feed contains information we need to process a feed.
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds reads and unmarshals the feed data file.
func RetrieveFeeds(dataFile string) ([]*Feed, error) {

	// Open file
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// Schedule the file to be closed once return
	defer file.Close()

	// Decode the file into a slice of Feed pointers
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	// Since a last row, we don't need to check err
	return feeds, err
}
