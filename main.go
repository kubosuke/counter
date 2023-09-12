package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	COUNTER_FILE = "counter.txt"
	LAYOUT       = "2006-01-02 15:04:05.999999999"
)

func counter(w http.ResponseWriter, r *http.Request) {
	// time
	time_current := time.Now()

	// Read file
	f, err := os.OpenFile(COUNTER_FILE, os.O_RDWR+os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")
	lines = lines[:len(lines)-1]

	// count last 1 min
	time_1min_ago := time_current.Add(-time.Second)
	timezone, _ := time.LoadLocation("Local")

	left := 0
	right := len(lines) - 1
	for left <= right {
		mid := (left + right) / 2
		t, err := time.ParseInLocation(LAYOUT, lines[mid], timezone)
		if err != nil {
			log.Fatal(err)
		}
		if t.After(time_1min_ago) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	result := len(lines) - left

	// Write file and response
	fmt.Fprintln(w, result)
	if _, err = f.WriteString(fmt.Sprintln(time_current.Format(LAYOUT))); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", counter)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
