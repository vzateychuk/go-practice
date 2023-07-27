package search

import "log"

// Result contains the result of a search.
type Result struct {
	Field   string
	Content string
}

// Matcher defines the behavior required by types that want to implement a new search type.
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match is launched as a goroutine for each individual feed to run searches concurrently.
func Match(matcher Matcher, feed *Feed, searchTert string, resultChan chan<- *Result) {
	//  Perform the search against the specified matcher.
	seachResults, err := matcher.Search(feed, searchTert)
	if err != nil {
		log.Println(err)
		return
	}

	// Write the results to the channel.
	for _, found := range seachResults {
		resultChan <- found
	}
}
