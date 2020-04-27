package main

import (
	"KExample/src/ppl"
	"fmt"
	"log"
	"net/http"
	"unicode"
)

var allPeoplle ppl.PeopleData

func handler(w http.ResponseWriter, r *http.Request) {
	requestName := r.URL.Path[1:]
	for index, char := range requestName {
		if index != 0 && unicode.IsUpper(char) {
			requestName = requestName[:index] + " " + requestName[index:]
		}
	}
	// fmt.Println("URL request name : " + r.URL.Path[1:] + ", processed name: " + requestName)
	for _, eachPpl := range allPeoplle.Peoples {
		if eachPpl.Name == requestName {
			fmt.Fprintf(w, "Quote for %s is: %s", requestName, eachPpl.Quote)
			return
		}
	}
	fmt.Fprintf(w, "Sorry, cannot find quote for '%s'", requestName)
	return
}

func main() {
	allPeoplle = ppl.InitFromFile()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
