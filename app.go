package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome!!</h1>")
    /* Multi-line tags and tags will be in the same format in the page source
	fmt.Fprintf(w,
		`<h1>Welcome!!</h1>
        <p>This is the first sample</p>
        <h2>Visit us again!!</h2>
    `)
    */
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page...")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.ListenAndServe(":8000", nil)
}
