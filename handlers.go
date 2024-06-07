package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"sync"
)

var (
	dataFile = "data.json"
	mutex    = &sync.Mutex{}
)

func addPerson(w http.ResponseWriter, r *http.Request) {

	var newPerson Person

	err := json.NewDecoder(r.Body).Decode(&newPerson)
	if err != nil {
		jsonResponse(w, Response{Success: false, Errors: "Invalid request body"}, http.StatusBadRequest)
		return
	}

	valid, validationError := validateIIN(newPerson.IIN)
	if !valid {
		jsonResponse(w, Response{Success: false, Errors: validationError}, http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	people, err := loadPeople()

	if err != nil {
		jsonResponse(w, Response{Success: false, Errors: "Failed to load data"}, http.StatusInternalServerError)
		return
	}

	for _, person := range people {
		if person.IIN == newPerson.IIN {
			jsonResponse(w, Response{Success: false, Errors: "Person with this IIN already exists"}, http.StatusBadRequest)
			return
		}
	}

	people = append(people, newPerson)

	err = savePeople(people)

	if err != nil {
		jsonResponse(w, Response{Success: false, Errors: "Failed to save data"}, http.StatusInternalServerError)
		return
	}

	jsonResponse(w, Response{Success: true}, http.StatusCreated)
}

func checkIin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	iin := vars["iin"]
	fmt.Println(iin)
	birth, gen, errorD := getBirthDateAndGender(iin)

	if errorD != nil {
		//jsonResponse(w, Response{Success: false, Errors: "Invalid request body"}, http.StatusBadRequest)
	}

	valid, validationError := validateIIN(iin)
	if !valid {
		jsonResponse(w, Response{Success: false, Errors: validationError}, http.StatusBadRequest)
		return
	}
	jsonResponse(w, CheckResponse{Correct: true, Sex: gen, DateOfBirth: birth}, http.StatusOK)
}

func GetInfoByIIN(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	iin := vars["iin"]

	valid, validationError := validateIIN(iin)
	if !valid {
		jsonResponse(w, Response{Success: false, Errors: validationError}, http.StatusInternalServerError)
		return
	}
	//jsonResponse(w, CheckResponse{Correct: true, Sex: gen, DateOfBirth: birth}, http.StatusOK)

	people, err := loadPeople()

	if err != nil {
		//jsonResponse(w, Response{Success: false, Errors: "Failed to load data"}, http.StatusInternalServerError)
		return
	}
	var person Person
	for _, p := range people {
		if p.IIN == iin {
			person = p
			break
		}
	}

	if person.IIN == "" {
		jsonResponse(w, Response{Success: false, Errors: "Не найдено"}, http.StatusNotFound)
		return
	}

	jsonResponse(w, Person{Name: person.Name, IIN: person.IIN, Phone: person.Phone}, http.StatusFound)
}

func GetInfoByFio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fio := vars["fio"]

	people, err := loadPeople()

	if err != nil {
		//jsonResponse(w, Response{Success: false, Errors: "Failed to load data"}, http.StatusInternalServerError)
		return
	}
	var person []Person
	fmt.Println("person", person)
	for _, p := range people {
		if strings.Contains(p.Name, fio) {
			person = append(person, p)
		}
	}
	if len(person) > 0 {
		jsonResponse(w, person, http.StatusFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	fmt.Fprintf(w, "[]")
	//jsonResponse(w, person, http.StatusNotFound)

}
