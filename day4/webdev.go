package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hey there")
	//you can even add varibles to your output stream
	fmt.Fprintf(w, " <p> this is how %s show variables %s </p>", "variables", "to the web")
	fmt.Fprintf(w, `this is a multiline parameterized functions 
		hey there`)
	// to organize the tabbing with the html templates later as usual

}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}
