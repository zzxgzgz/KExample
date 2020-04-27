package main

import (
	"KExample/src/ppl"
	"fmt"
	"log"
	"net/http"
	"unicode"
)

var peopleMap map[string]string

func handler(w http.ResponseWriter, r *http.Request) {
	requestName := r.URL.Path[1:]
	for index, char := range requestName {
		if index != 0 && unicode.IsUpper(char) {
			requestName = requestName[:index] + " " + requestName[index:]
		}
	}
	fmt.Println("URL request name : " + r.URL.Path[1:] + ", processed name: " + requestName)
	if quote, ok := peopleMap[requestName]; ok {
		fmt.Fprintf(w, "Quote for %s is: %s", requestName, quote)
	} else {
		fmt.Fprintf(w, "Sorry, cannot find quote for '%s'", requestName)
	}
	return
}

func main() {
	peopleMap = ppl.InitFromFile()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
