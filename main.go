package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	pass := "12345"
	hashedPass, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}

	err = comparePassword(pass, hashedPass)
	if err != nil {
		log.Fatal("Not logged in")
	}
	log.Println("Logged in!")
}
func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error while generating bcrypt hash from password: %w", err)
	}
	return bs, nil
}

func comparePassword(password string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), hashedPass)
	if err != nil {
		return fmt.Errorf("Invalid password: %w", err)
	}
	return nil
}
