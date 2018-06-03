package main

import (
	"log"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://nleiva.github.io",
		http.StatusTemporaryRedirect)
}

func main() {
	http.HandleFunc("/", redirect)
	http.HandleFunc("/no-redirect", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello There!"))
	})
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
