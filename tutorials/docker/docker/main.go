package main

import (
	"log"
	"net/http"
	"fmt"
)



func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "hello world\n")
	})
	log.Print("Running server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}