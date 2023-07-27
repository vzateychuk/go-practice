package search

import "log"

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Register is called to register a matcher for use by the program.
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}

// defaultMatcher implements the default matcher.
type defaultMatcher struct{}

// Unlike when you call methods directly from values and pointers, when you call a method via an interface type value, the rules are different.
// Methods declared with pointer receivers can only be called by interface type values that contain pointers.
// Methods declared with value receivers can be called by interface type values that contain both values and pointers.
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {

	// Search implements the behavior for the default matcherFatalln(v ...any)
	return nil, nil
}
