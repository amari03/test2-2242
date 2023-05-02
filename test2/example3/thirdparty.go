/*
Using third party packages instead of creating your own middleware.
third-party middleware we'll demonstrate is goji/httpauth, 
which provides HTTP Basic Authentication functionality.
*/
package main

import (
	"log"
	"net/http"

	"github.com/goji/httpauth"
)

func main() {
	authHandler := httpauth.SimpleBasicAuth("addie", "pa$$word")

	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", authHandler(finalHandler))

	log.Print("Listening on :4000...")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Working!"))
}