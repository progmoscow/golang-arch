package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Name string
	Age  int
}

func main() {

	http.HandleFunc("/encode", foo)
	http.ListenAndServe(":8000", nil)
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
