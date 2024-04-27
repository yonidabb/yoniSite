package main

import (
	"fmt"
	"log"
	"net/http"
)

// ... your existing handler functions
// https://www.free-css.com/free-css-templates/page292/microo
func main() {
	// Serve static files from the "public" directory
	fs := http.FileServer(http.Dir("./microo-html"))
	http.Handle("/", http.StripPrefix("/", fs))

	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
