package ppl

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type PeopleList struct {
	Data []People `json:"data"`
}

type People struct {
	Name  string `json:"name"`
	Quote string `json:"quote"`
}

func InitFromFile() map[string]string {
	var allPeople PeopleList

	absPath, _ := filepath.Abs("../ppl/data.json")

	// fmt.Println("File path: ", absPath)

	jsonFile, err := os.Open(absPath)

	defer jsonFile.Close()

	if err != nil {
		panic(err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &allPeople)

	peopleMap := make(map[string]string)

	for _, eachPeople := range allPeople.Data {
		peopleMap[eachPeople.Name] = eachPeople.Quote
		// fmt.Println("Name: " + eachPeople.Name + ", quote: " + eachPeople.Quote)
	}

	return peopleMap
}
