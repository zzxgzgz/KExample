package main

import (
	"KExample/src/ppl"
	"fmt"
	"log"
	"net/http"
)

var allPeoplle ppl.PeopleData

func handler(w http.ResponseWriter, r *http.Request) {
	var peoplesQuote = "We cannot find it..."

	for _, eachPpl := range allPeoplle.Peoples {
		if eachPpl.Name == r.URL.Path[1:] {
			peoplesQuote = eachPpl.Quote
			break
		}
	}
	fmt.Fprintf(w, "Quote for %s is: %s", r.URL.Path[1:], peoplesQuote)
}

func main() {
	allPeoplle = ppl.InitFromFile()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
