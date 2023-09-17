package BL

import (
	"encoding/json"
	"log"
)

type Person struct {
	id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

func ToJSON(p Person) bytes {
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return jsonBytes
}

func FromJSON(s bytes) Person {
	person := Person{}
	err := json.Unmarshal(s, &person)
	if err != nil {
		log.Fatal(err)
		return Person{}
	}

	return person
}
