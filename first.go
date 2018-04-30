package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"fmt"
	"encoding/json"
)


type Person struct {
	ID int `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Address string `json:"address,omitempty"`
}

var people []Person

func main() {

	people = append(people, Person{ID: 1, Firstname: "Sujal", Lastname: "Patel", Address: "IN"})
	people = append(people, Person{ID: 2, Firstname: "Tom", Lastname: "Jone", Address: "US"})

	route := mux.NewRouter();

	route.HandleFunc("/people", GetPeoples).Methods("GET");
	route.HandleFunc("/people/{id}", People).Methods("GET");
	route.HandleFunc("/people/{id}", CreatePeople).Methods("POST");
	route.HandleFunc("/people/{id}", DeletePeople).Methods("DELETE");

/*	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprintf(writer, "Hello, %q", html.EscapeString(request.URL.Path))
	})
*/
	log.Fatal(http.ListenAndServe(":2209", route))

}

func GetPeoples(writer http.ResponseWriter, request *http.Request)  {
	json.NewEncoder(writer).Encode(people[0].Firstname + people[0].Lastname)
}

func People(writer http.ResponseWriter, request *http.Request)  {
	fmt.Fprintf(writer, "People/id GET Route");
}

func CreatePeople(writer http.ResponseWriter, request *http.Request)  {
	fmt.Fprintf(writer, "People/id POST Route");
}

func DeletePeople(writer http.ResponseWriter, request *http.Request)  {
	fmt.Fprintf(writer, "People/id DELETE Route");
}