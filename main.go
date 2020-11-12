package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	Name string
	Age  int
}

func main() {

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8000", nil)
}

func bar(w http.ResponseWriter, r *http.Request) {
	people := []person{}
	err := json.NewDecoder(r.Body).Decode(&people)

	if err != nil {
		log.Println("Bad data to decode", err)
	}
	fmt.Println(people)
}

func foo(w http.ResponseWriter, r *http.Request) {
	person01 := person{
		Name: "Tom",
		Age:  10,
	}
	person02 := person{
		Name: "Dave",
		Age:  19,
	}

	people := []person{person01, person02}

	err := json.NewEncoder(w).Encode(people)

	if err != nil {
		log.Println("Bad data to encode ", err)
	}
}
