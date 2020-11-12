package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
	Admin bool
}

func main() {
	// p1 := person{
	// 	First: "Tom",
	// }
	// p2 := person{
	// 	First: "James",
	// }
	// xp := []person{p1, p2}

	// bs, err := json.Marshal(xp)

	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println("PRINT JSON", string(bs))
	// xp2 := []person{}
	// err = json.Unmarshal(bs, &xp2)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println("back into a Go data structure", xp2)

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	person01 := person{
		Name : "Tom",
		Age : 20
	}

	err := json.NewEncoder(w).Encode(person01)
	if err != nil {
		log.Println("Encode bad data", err)	
	}
}

func bar(w http.ResponseWriter, r *http.Request) {

}
