package bts

import (
	"fmt"
	"reflect"
)

type ByteWriter struct {
	bytes []byte
}

func (b *ByteWriter) WriteString(str string) {
	b.bytes = append(b.bytes, byte(len(str)))
	b.bytes = append(b.bytes, []byte(str)...)
}

func (b *ByteWriter) Bytes() []byte {
	return b.bytes
}

func (b *ByteWriter) Write(t any) {
	switch v := t.(type) {
	case *string:
		b.WriteString(*v)
	default:
		fmt.Println("Nope", t)
	}
}

func Encode(m any) []byte {
	var bw ByteWriter

	v := reflect.ValueOf(m).Elem()

	for i := 0; i < v.NumField(); i++ {
		bw.Write(v.Field(i).Addr().Interface())
	}

	return bw.Bytes()
}
