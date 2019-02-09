package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty`
	Address   *Address `json:"address,omitempty`
}
type Address struct {
	City  string `json:city,omitempty`
	State string `json:state,omitempty`
}

var people []Person

func main() {
	router := mux.NewRouter()

	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City x", State: "State x"}})
	people = append(people, Person{ID: "2", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City x", State: "State x"}})
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City x", State: "State x"}})

	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	//	log.Logger
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request)    {}
func GetPerson(w http.ResponseWriter, r *http.Request)    {}
func CreatePerson(w http.ResponseWriter, r *http.Request) {}
func DeletePerson(w http.ResponseWriter, r *http.Request) {}
