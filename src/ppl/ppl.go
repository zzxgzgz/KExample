package ppl

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"fmt"
)

type PeopleList struct {
	Data []People `json:"data"`
}

type People struct {
	Name  string `json:"name"`
	Quote string `json:"quote"`
}

//	InitFromFile reads the data.json file and constructs the map, the key of which is a person's name and the value of that key is a quote for that person
func InitFromFile() map[string]string {
	var allPeople PeopleList

	absPath, _ := filepath.Abs("/etc/KExample/src/ppl/data.json")
	pwd, _ := os.Getwd();
	fmt.Println("Abs path: "+ absPath)
	fmt.Println("PWD: " + pwd)
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
	}

	return peopleMap
}
