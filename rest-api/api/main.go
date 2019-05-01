package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Person struct {
	FirstName string  `json:"fname"`
	LastName  string  `json:"lname"`
	Age       int     `json:"age"`
	Addr      Address `json:"address"`
}

type Address struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

var persons []Person

func getPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(persons)

}

func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, val := range persons {
		if val.FirstName == params["fname"] {
			_ = json.NewEncoder(w).Encode(val)
			return
		}
	}
	//json.NewEncoder(w).Encode(Person{})
}

func createNewPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person Person
	_ = json.NewDecoder(r.Body).Decode(person)
	persons = append(persons, person)
	json.NewEncoder(w).Encode(person)

}

func main() {
	router := mux.NewRouter()
	p1 := Person{"Rohit", "Bhanot", 30, Address{"Gurgaon", "Haryana", "India"}}
	p2 := Person{FirstName: "Mohit", LastName: "Bhanot", Age: 30, Addr: Address{"Mumbai", "MH", "India"}}
	persons = append(persons, p1, p2)
	router.HandleFunc("/api/person/all", getPersons).Methods("GET")
	router.HandleFunc("/api/person/{fname}", getPerson).Methods("GET")
	router.HandleFunc("/api/person/new", createNewPerson).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}
