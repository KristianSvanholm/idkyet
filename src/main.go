package main

import (
	"fmt"
	bts "serialization/internal/bytes"
)

type Extended struct {
	Email    string
	Password string
	Age      int8
	A        int16
	B        int32
	C        int64
	D        float32
	Rot      float64
}

func main() {

	//b := Base{id: 123, name: "hello"}
	v := Extended{Email: "krs@mail.com", Password: "verySecret", Age: 24, A: 9999, B: 999999999, C: 999999999999999999, D: 34.6, Rot: 0.30000000000000004}

	bytes := bts.Encode(&v)
	var w Extended

	bts.Decode(&w, bytes)

	fmt.Println("Original: ", v)
	fmt.Println("Encoded:", bytes)
	fmt.Println("Decoded: ", w)
}
