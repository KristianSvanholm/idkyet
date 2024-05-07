package bts

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
)

type ByteWriter struct {
	bytes []byte
}

func (b *ByteWriter) WriteString(str string) {
	b.bytes = append(b.bytes, byte(len(str)))
	b.bytes = append(b.bytes, []byte(str)...)
}

func (b *ByteWriter) WriteInt(i int) {
	b.bytes = append(b.bytes, byte(i))
}

func (b *ByteWriter) WriteFloat64(f float64) {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))

	b.bytes = append(b.bytes, buf[:]...)
}

func (b *ByteWriter) Bytes() []byte {
	return b.bytes
}

func (b *ByteWriter) Write(t any) {
	switch v := t.(type) {
	case *string:
		b.WriteString(*v)
	case *int:
		b.WriteInt(*v)
	case *float64:
		b.WriteFloat64(*v)
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
