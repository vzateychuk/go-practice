package main

import (
	"fmt"
	"net/http"
	"time"
	"vez/backpressure/back"
)

func delayFunction(waitSec int) string {
	time.Sleep(time.Duration(waitSec) * time.Second)
	return fmt.Sprintf("done waiting: %d sec", waitSec)
}

func main() {

	bp := back.NewBackpressure(2)
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		err := bp.Process(func() {
			w.Write([]byte(delayFunction(5)))
			w.WriteHeader(http.StatusOK)
		})

		if err != nil {
			w.Write([]byte("Too many requests"))
			w.WriteHeader(http.StatusTooManyRequests)
		}
	}

	http.HandleFunc("/api", handleFunc)
	http.ListenAndServe(":8080", nil)

}
