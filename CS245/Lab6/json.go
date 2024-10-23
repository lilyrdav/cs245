package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

type Birth struct {
	Year int
	Month string
	Day int
}

type Student struct {
	FirstName string
	LastName string
	ID string
	Age int
	DateOfBirth Birth
	FavoriteRealNumber int
}

type jsonMap map[string]interface{}
type jsonSlice []jsonMap

func main() {
	jsonFile, err := os.Open("xxx.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened xxx.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var students jsonSlice
    json.Unmarshal(byteValue, &students)    

	jsonFile.Close()

	fmt.Println(students)

	data, _ := json.Marshal(students)
	    fmt.Println(string(data))
}
