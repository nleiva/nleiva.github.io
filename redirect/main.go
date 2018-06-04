package main

import (
	"crypto/rand"
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/net/http2"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	// StatusMovedPermanently  = 301
	// StatusTemporaryRedirect = 307
	http.Redirect(w, r, "https://nleiva.github.io",
		http.StatusMovedPermanently)
}

func main() {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("www.nleiva.com", "nleiva.com"), //Your domain here
		Cache:      autocert.DirCache("certs"),                             //Folder for storing certificates
	}
	// Routing
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/redirect.html")
	})
	http.HandleFunc("/redirect", redirect)

	tlsConfig := &tls.Config{
		Rand:           rand.Reader,
		Time:           time.Now,
		NextProtos:     []string{http2.NextProtoTLS, "http/1.1"},
		MinVersion:     tls.VersionTLS12,
		GetCertificate: certManager.GetCertificate,
	}

	server := &http.Server{
		Addr:      ":https",
		TLSConfig: tlsConfig,
	}

	// If fallback is nil, the returned handler redirects all GET and HEAD requests to
	// the default TLS port 443 with 302 Found status code, preserving the original request
	// path and query.
	go http.ListenAndServe(":http", certManager.HTTPHandler(nil))

	log.Fatal(server.ListenAndServeTLS("", "")) //Key and cert are coming from Let's Encrypt

	//err := http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/www.nleiva.com/fullchain.pem",
	//	"/etc/letsencrypt/live/www.nleiva.com/privkey.pem", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
}
