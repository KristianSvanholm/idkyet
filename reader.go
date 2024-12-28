package idkyet

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"reflect"
)

type ByteReader struct {
	bytes       []byte
	pos         int
	outOfBounds bool
	count       int
}

func NewReader(bytes []byte) ByteReader {
	return ByteReader{bytes: bytes, count: len(bytes)}
}

func (b *ByteReader) r() byte {
	if b.pos >= b.count {
		b.outOfBounds = true
		return 0x0
	}
	v := b.bytes[b.pos]
	b.pos += 1
	return v
}

func (b *ByteReader) rr(len int) []byte {
	if b.pos+len > b.count {
		b.outOfBounds = true
		return empty(len)
	}
	v := b.bytes[b.pos : b.pos+len]
	b.pos += len
	return v
}

// Retuens n long zeroed bytes buffer
func empty(n int) []byte {
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = 0x0
	}
	return bytes
}

// TODO:: Convert to use endOfString value to prevent i64 size bloat
func (b *ByteReader) ReadString() string {
	return string(b.rr(int(b.r())))
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

func (b *ByteReader) ReadStruct(v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		b.Read(v.Field(i))

	}
}

func (b *ByteReader) ReadArray(v reflect.Value) {
	for i := 0; i < v.Type().Len(); i++ {
		b.Read(v.Index(i))
	}
}

func (b *ByteReader) Bytes() []byte {
	return b.bytes
}

func (b *ByteReader) Read(v reflect.Value) {

	nv := reflect.Indirect(reflect.New(v.Type()))

	switch v.Kind() {
	case reflect.String:
		nv.SetString(b.ReadString())
	case reflect.Int8:
		nv.SetInt(int64(b.ReadInt8()))
	case reflect.Int16:
		nv.SetInt(int64(b.ReadInt16()))
	case reflect.Int32:
		nv.SetInt(int64(b.ReadInt32()))
	case reflect.Int64:
		nv.SetInt(b.ReadInt64())
	case reflect.Float32:
		nv.SetFloat(float64(b.ReadFloat32()))
	case reflect.Float64:
		nv.SetFloat(b.ReadFloat64())
	case reflect.Struct:
		b.ReadStruct(nv)
	case reflect.Array:
		b.ReadArray(nv)
	default:
		fmt.Println("Nope Read", v)
	}

	v.Set(nv)
}

func Decode(m any, b []byte) error {
	br := NewReader(b)

	v := reflect.ValueOf(m).Elem()
	br.Read(v)

	if br.outOfBounds {
		return errors.New("Attempted to read outside bytes buffer. Some fields may be empty.")
	}

	return nil
}
