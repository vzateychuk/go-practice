package main

import (
	"log"
	"os"
	"strings"

	_ "vez/puller/matchers"
	"vez/puller/search"
)

func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

func main() {

	log.Println("Environment:")
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		log.Printf("[%s] = [%s] \n", pair[0], pair[1])
	}

	log.Println("Started")
	// Perform the search for the specified term
	search.Run("president")
}
