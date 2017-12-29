package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "whoa, go is ready to meat")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "expert web design by aloy aditya sen")
}
func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000", nil)

}
