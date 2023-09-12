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
	var idx = 0
	timezone, _ := time.LoadLocation("Local")
	for i, line := range lines {
		idx = i
		t, err := time.ParseInLocation(LAYOUT, line, timezone)
		if err != nil {
			log.Fatal(err)
		}
		if t.After(time_1min_ago) {
			break
		}
	}
	result := len(lines) - idx - 1

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
