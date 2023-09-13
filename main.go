package main

import (
	"log"
	"net/http"

	"counter/handler"
)

func main() {
	http.HandleFunc("/", handler.CounterHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
