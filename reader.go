package idkyet

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
)

type ByteReader struct {
	bytes []byte
	pos   int
}

func NewReader(bytes []byte) ByteReader {
	return ByteReader{bytes: bytes}
}

func (b *ByteReader) r() byte {
	v := b.bytes[b.pos]
	b.pos += 1
	return v
}

func (b *ByteReader) rr(len int) []byte {
	v := b.bytes[b.pos : b.pos+len]
	b.pos += len
	return v
}

// TODO:: Convert to use endOfString value to prevent i64 size bloat
func (b *ByteReader) ReadString() string {
	len := int(b.r())
	str := string(b.rr(len))

	return str
}

// TODO:: Might be able to consolidate some logic here for the int and float variations
func (b *ByteReader) ReadInt8() int8 {
	return int8(b.r())
}

func (b *ByteReader) ReadInt16() int16 {
	return int16(binary.LittleEndian.Uint16(b.rr(2)))
}

func (b *ByteReader) ReadInt32() int32 {
	return int32(binary.LittleEndian.Uint32(b.rr(4)))
}

func (b *ByteReader) ReadInt64() int64 {
	return int64(binary.LittleEndian.Uint64(b.rr(8)))
}

func (b *ByteReader) ReadFloat32() float32 {
	bits := binary.LittleEndian.Uint32(b.rr(4))
	return math.Float32frombits(bits)
}

func (b *ByteReader) ReadFloat64() float64 {
	bits := binary.LittleEndian.Uint64(b.rr(8))
	return math.Float64frombits(bits)
}

func (b *ByteReader) Bytes() []byte {
	return b.bytes
}

func (b *ByteReader) Read(t any) {
	switch v := t.(type) {
	case *string:
		*v = b.ReadString()
	case *int8:
		*v = b.ReadInt8()
	case *int16:
		*v = b.ReadInt16()
	case *int32:
		*v = b.ReadInt32()
	case *int64:
		*v = b.ReadInt64()
	case *float32:
		*v = b.ReadFloat32()
	case *float64:
		*v = b.ReadFloat64()
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
