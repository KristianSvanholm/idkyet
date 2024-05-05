package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type ByteWriter struct {
	bytes []byte
}

func (b *ByteWriter) writeString(str string) {
	b.bytes = append(b.bytes, byte(len(str)))
	b.bytes = append(b.bytes, []byte(str)...)
}

func (e *Extended) serialize() []byte {
	var bw ByteWriter

	bw.writeString(e.Password)
	bw.writeString(e.Email)

	return bw.bytes
}

type ByteReader struct {
	bytes    []byte
	position int
}

func (b *ByteReader) readString() string {
	len := int(b.bytes[b.position])
	b.position += 1

	str := string(b.bytes[b.position : b.position+len])

	b.position += len

	return str
}

func (e *Extended) deserialize(b []byte) {
	br := ByteReader{bytes: b}

	e.Password = br.readString()
	e.Email = br.readString()
}

type Extended struct {
	Email    string
	Password string
}

func main() {

	//b := Base{id: 123, name: "hello"}
	v := Extended{Email: "krs@mail.com", Password: "frick"}
	//fmt.Println(v)
	bytesEnc := v.serialize()
	bytesJson, _ := json.Marshal(v)
	fmt.Println(len(bytesEnc), " - ", bytesEnc)
	fmt.Println(len(bytesJson), " - ", bytesJson)

	fmt.Println(string(bytesEnc))
	fmt.Println(string(bytesJson))
	var nv Extended
	nv.deserialize(bytesEnc)
	//fmt.Println(nv)

	//x := GetSqlColumnToFieldMap(&v)
	//fmt.Println(x)
}

// Converts struct into slice of values existing in DB using reflection.
func GetSqlColumnToFieldMap(model any) []any {
	fmt.Println(reflect.TypeOf(model))
	//t := reflect.TypeOf(model).Elem()
	v := reflect.ValueOf(model).Elem() // Value of model
	s := make([]any, 0)
	for i := 0; i < v.NumField(); i++ { // Loop over fields
		var val any
		if v.Field(i).Kind() == reflect.Struct {
			fmt.Println("hello!")
			//tt := t.Field(i).Type
			fmt.Println(v.Field(i).Interface())
			val = GetSqlColumnToFieldMap(v.Field(i))
		} else {
			val = v.Field(i).Addr().Interface()
		}
		s = append(s, val) // Add value to slice
	}
	return s
}
