package main

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/people/info", addPerson).Methods("POST")
	r.HandleFunc("/iin_check/{iin}", checkIin).Methods("GET")
	r.HandleFunc("/people/info/iin/{iin}", GetInfoByIIN).Methods("GET")
	r.HandleFunc("/people/info/fio/{fio}", GetInfoByFio).Methods("GET")
}
