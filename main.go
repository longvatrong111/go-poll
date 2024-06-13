package main

import (
	"net/http"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Welcome to my website!")
	// })

	fs := http.FileServer(http.Dir("./pollsite/build"))
	// fs := http.FileServer(http.Dir("static/"))

	http.Handle("/", fs)
	http.ListenAndServe(":80", nil)
}
