package main

import (
	"log"
	"net/http"
)

// HelloServer is an example from
// https://gist.github.com/denji/12b3a568f092ab951456
func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/www.nleiva.com/fullchain.pem",
		"/etc/letsencrypt/live/www.nleiva.com/privkey.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
