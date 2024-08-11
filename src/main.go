package main

import (
	"fmt"
	bts "serialization/internal/bytes"
)

type Account struct {
	Email string
	Pwd   string
	Age   int8
	Bal   int64
}

func main() {

	v := Account{Email: "krs@mail.com", Pwd: "verySecret", Age: 24, Bal: 25600}

	bytes := bts.Encode(&v)
	var w Account
	bts.Decode(&w, bytes)

	fmt.Println("Original: ", v)
	fmt.Println("Encoded:", len(bytes), bytes)
	fmt.Println("Decoded: ", w)
}
