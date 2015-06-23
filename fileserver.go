package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", send)
	http.ListenAndServe(":4000", nil)
}

func send(w http.ResponseWriter, r *http.Request) {
	// send a gopher!
	http.ServeFile(w, r, "gopher.jpg")
}
