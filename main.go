package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	RegisterRoutes(r)

	http.Handle("/", r)
	fmt.Println("Сервер запущен на порту 8000")
	http.ListenAndServe(":8000", nil)
}
