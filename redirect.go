package main

import (
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://www.uni-trier.de" + r.URL.String(), http.StatusMovedPermanently)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
