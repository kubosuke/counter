package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	COUNTER_FILE = "counter.txt"
	LAYOUT       = "2006-01-02 15:04:05.999999999"
)

var (
	data []string
	mu   sync.Mutex
)

func init() {
	f, err := os.OpenFile(COUNTER_FILE, os.O_RDONLY+os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")
	data = lines[:len(lines)-1]
}

func CounterHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	time_current := time.Now()

	time_1min_ago := time_current.Add(-time.Minute)
	idx := 0
	timezone, _ := time.LoadLocation("Local")
	for i, line := range data {
		idx = i
		t, err := time.ParseInLocation(LAYOUT, line, timezone)
		if err != nil {
			log.Fatal(err)
		}
		if t.After(time_1min_ago) {
			break
		}
	}
	result := len(data) - idx

	fmt.Fprintln(w, result)
	data = append(data, fmt.Sprint(time_current.Format(LAYOUT)))
	f, err := os.OpenFile(COUNTER_FILE, os.O_APPEND+os.O_WRONLY+os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintln(time_current.Format(LAYOUT))); err != nil {
		log.Fatal(err)
	}
}
