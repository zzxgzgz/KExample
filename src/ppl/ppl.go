package ppl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type PeopleData struct {
	Peoples []People `json:"data"`
}

type People struct {
	Name  string `json:"name"`
	Quote string `json:"quote"`
}

func InitFromFile() PeopleData {
	var allPeople PeopleData

	absPath, _ := filepath.Abs("../ppl/data.json")

	fmt.Println("File path: ", absPath)

	jsonFile, err := os.Open(absPath)

	defer jsonFile.Close()

	if err != nil {
		panic(err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &allPeople)

	fmt.Println("All people in the data set: ")

	for i := 0; i < len(allPeople.Peoples); i++ {
		fmt.Println("Name: ", allPeople.Peoples[i].Name)
		fmt.Println("Quote: ", allPeople.Peoples[i].Quote)
	}
	return allPeople
}
