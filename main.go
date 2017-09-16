package main

import (
	"fmt"
	"log"
	"net/http"
)

var version = "dev"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world from example %s!\n", version)
	})
	log.Println("listening on port 5000...")
	if err := http.ListenAndServe(":5000", mux); err != nil {
		log.Fatalln(err)
	}
}
