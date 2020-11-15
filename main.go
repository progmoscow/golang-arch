package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var key = []byte{}}

type person struct {
	Name string
	Age  int
}

func main() {

	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}
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

func signMessage(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)
	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("Error in signMessage while hashing message: %w",err)
	}

	signature := h.Sum(nil)
	return signature, nil
}

func checkSignature(msg ,sig []byte)(bool,error){
	newSign := signMessage(msg)
	if err != nil {
		return false,fmt.Errorf("Error in checkSign while getting signature of message: %w", err)
	}
	same := hmac.Equal(newSign,sig)
	return same, nil
}
 