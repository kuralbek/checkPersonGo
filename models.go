package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	IIN   string `json:"iin"`
	Phone string `json:"phone"`
}

type Response struct {
	Success bool   `json:"success"`
	Errors  string `json:"errors,omitempty"`
}

type CheckResponse struct {
	Correct     bool   `json:"Correct"`
	Sex         string `json:"Sex"`
	DateOfBirth string `json:"DateOfBirth"`
}

func loadPeople() ([]Person, error) {
	var people []Person

	file, err := os.ReadFile(dataFile)

	if err != nil {
		if os.IsNotExist(err) {
			return people, nil
		}
		return nil, err
	}

	err = json.Unmarshal(file, &people)
	if err != nil {

		return nil, err
	}

	return people, nil
}

func savePeople(people []Person) error {
	data, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dataFile, data, 0644)
}
