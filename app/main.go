package main

import (
	"fmt"
	"net/http"
	"log"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have visited %s!", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handle)
	fmt.Println("Starting web server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}