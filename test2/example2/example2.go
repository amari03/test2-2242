/*
building a service which processes requests containing a JSON body.
a) checks for the existence of a Content-Type header
b) if the header exists, check that it has the mime type application/json.

*/
package main

import (
	"log"
	"mime"
	"net/http"
)

//crerating the middleware function
func enforceJSONHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		//this is the error created if the checks failed
		if contentType != "" {
			mt, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
				return
			}

			if mt != "application/json" {
				http.Error(w, "Content-Type header must be application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ALL GOOD!"))
}

func main() {
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", enforceJSONHandler(finalHandler))

	log.Print("Listening on :4000...")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}	

/*
making some requests using cURL:
curl -i localhost:4000
curl -i -H "Content-Type: application/xml" localhost:4000
curl -i -H "Content-Type: application/json; charset=UTF-8" localhost:4000
*/