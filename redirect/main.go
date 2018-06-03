package main

import (
	"log"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://nleiva.github.io", 301)
}

func main() {
	http.HandleFunc("/", redirect)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
