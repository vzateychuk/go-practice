package search

import (
	"fmt"
	"log"
	"sync"
)

// A map of registered matchers for searching.
var matchers = make(map[string]Matcher)

func Run(searchTerm string) {

	// Retrieve list of feeds
	feeds, err := RetrieveFeeds("data/data.json")
	if err != nil {
		log.Fatal(err)
	}

	// Create unbuffred channel to receive match results
	results := make(chan *Result)

	// Setup a wait group so we can process all feeds
	var wg sync.WaitGroup

	// The number of goroutines we need
	wg.Add(len(feeds))

	// Launch a goroutine for each feed to find the result
	for _, feed := range feeds {
		// Retrieve a matcher for the search
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// Launch goroutine to perform the search
		go func(m Matcher, fd *Feed) {
			Match(m, fd, searchTerm, results)
			wg.Done()
		}(matcher, feed)
	}

	// Finally launch goroutine to monitor when all the work is done
	go func() {
		// Waitfor everything is complete
		wg.Wait()

		// Close channel to signal to Disply
		close(results)
	}()

	// Start displaying result as they are available
	Display(results)
}

// Display writes results to the terminal window as they  are received by the individual goroutines
func Display(results chan *Result) {

	// The channel blocks until a result is written to the channel.
	// Once the channel is closed the for loop terminates.
	for res := range results {
		fmt.Printf("%s:\n%s\n\n", res.Field, res.Content)
	}
}
