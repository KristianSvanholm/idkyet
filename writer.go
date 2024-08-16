package idkyet

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
)

type ByteWriter struct {
	bytes []byte
}

func (b *ByteWriter) w(bs ...byte) {
	b.bytes = append(b.bytes, bs...)
}

func (b *ByteWriter) WriteString(str string) {
	b.w(byte(len(str)))
	b.w([]byte(str)...)
}

// TODO:: Might be able to consolidate some logic here for the int and float variations
func (b *ByteWriter) WriteInt8(i int8) {
	b.w(byte(i))
}

func (b *ByteWriter) WriteInt16(i int16) {
	var buf [2]byte
	binary.LittleEndian.PutUint16(buf[:], uint16(i))

	b.w(buf[:]...)
}

func (b *ByteWriter) WriteInt32(i int32) {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], uint32(i))

	b.w(buf[:]...)
}

func (b *ByteWriter) WriteInt64(i int64) {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], uint64(i))

	b.w(buf[:]...)
}

func (b *ByteWriter) WriteFloat32(f float32) {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], math.Float32bits(f))

	b.w(buf[:]...)
}

func (b *ByteWriter) WriteFloat64(f float64) {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], math.Float64bits(f))

	b.w(buf[:]...)
}

func (b *ByteWriter) Bytes() []byte {
	return b.bytes
}

func (b *ByteWriter) Write(t any) {
	switch v := t.(type) {
	case *string:
		b.WriteString(*v)
	case *int8:
		b.WriteInt8(*v)
	case *int16:
		b.WriteInt16(*v)
	case *int32:
		b.WriteInt32(*v)
	case *int64:
		b.WriteInt64(*v)
	case *float32:
		b.WriteFloat32(*v)
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
