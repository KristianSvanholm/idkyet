package main

import (
	"fmt"
	bts "serialization/internal/bytes"
)

type Extended struct {
	Email    string
	Password string
	Age      int
	Rot      float64
}

func main() {

	//b := Base{id: 123, name: "hello"}
	v := Extended{Email: "krs@mail.com", Password: "frick", Age: 24, Rot: 34.6}

	bytes := bts.Encode(&v)
	var w Extended
	bts.Decode(&w, bytes)

	fmt.Println("Original: ", v)
	fmt.Println("Encoded:", bytes)
	fmt.Println("Decoded: ", w)
}
