package main

import (
	"encoding/base64"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))
}
