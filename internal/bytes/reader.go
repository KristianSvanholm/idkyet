package bts

import (
	"fmt"
	"reflect"
)

type ByteReader struct {
	bytes    []byte
	position int
}

func NewReader(bytes []byte) ByteReader {
	return ByteReader{bytes: bytes}
}

func (b *ByteReader) ReadString() string {
	len := int(b.bytes[b.position])
	b.position += 1

	str := string(b.bytes[b.position : b.position+len])

	b.position += len

	return str
}

func (b *ByteReader) ReadInt() int {
	i := int(b.bytes[b.position])
	b.position += 1
	return i
}

func (b *ByteReader) Bytes() []byte {
	return b.bytes
}

func (b *ByteReader) Read(t any) {
	switch v := t.(type) {
	case *string:
		*v = b.ReadString()
	case *int:
		*v = b.ReadInt()
	default:
		fmt.Println("Nope", t)
	}
}

func Decode(m any, b []byte) {
	br := NewReader(b)

	v := reflect.ValueOf(m).Elem()

	for i := 0; i < v.NumField(); i++ {
		br.Read(v.Field(i).Addr().Interface())
	}
}
