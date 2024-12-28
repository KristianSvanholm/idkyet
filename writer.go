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

func (b *ByteWriter) WriteStruct(v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		b.Write(v.Field(i))
	}
}

func (b *ByteWriter) WriteArray(v reflect.Value) {
	for i := 0; i < v.Type().Len(); i++ {
		b.Write(v.Index(i))
	}
}

func (b *ByteWriter) Bytes() []byte {
	return b.bytes
}

func (b *ByteWriter) Write(t reflect.Value) {
	switch t.Kind() {
	case reflect.String:
		b.WriteString(t.String())
	case reflect.Int8:
		b.WriteInt8(int8(t.Int()))
	case reflect.Int16:
		b.WriteInt16(int16(t.Int()))
	case reflect.Int32:
		b.WriteInt32(int32(t.Int()))
	case reflect.Int64:
		b.WriteInt64(t.Int())
	case reflect.Float32:
		b.WriteFloat32(float32(t.Float()))
	case reflect.Float64:
		b.WriteFloat64(t.Float())
	case reflect.Struct:
		b.WriteStruct(t)
	case reflect.Array:
		b.WriteArray(t)
	default:
		fmt.Println("Nope Write", t)
	}
}

func Encode(m any) []byte {
	var bw ByteWriter

	v := reflect.ValueOf(m).Elem()

	bw.Write(v)

	return bw.Bytes()
}
