package main

import (
	"KExample/src/ppl"
	"fmt"
	"log"
	"net/http"
	"unicode"
)

//	peopleMap stores the names and quotes of the data.json file.
var peopleMap map[string]string

// hander handler function that processes the request, checks if that name exists in our data and writes the result.
func handler(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query()["name"]) < 1 {
		fmt.Fprintf(w, "Please enter something.\n")
		return
	}
	requestName := r.URL.Query()["name"][0]
	for index, char := range requestName {
		if index != 0 && unicode.IsUpper(char) {
			requestName = requestName[:index] + " " + requestName[index:]
		}
	}
	// fmt.Println("URL request name : " + r.URL.Path[1:] + ", processed name: " + requestName)
	if quote, ok := peopleMap[requestName]; ok {
		fmt.Fprintf(w, "Quote for %s is: %s\n", requestName, quote)
	} else {
		fmt.Fprintf(w, "Sorry, cannot find quote for '%s'\n", requestName)
	}
	return
}

//	main main function that initiates the server
func main() {
	peopleMap = ppl.InitFromFile()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
