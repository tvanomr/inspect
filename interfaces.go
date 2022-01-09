package inspect

import (
	"encoding"
	"encoding/gob"
	"io"
)

type Reader interface {
	SetReader(io.Reader)
	Int32() (int32, error)
	Int64() (int64, error)
	Float32() (float32, error)
	Float64() (float64, error)
	String() (string, error)
	Bytes() ([]byte, error)
	ByteString() ([]byte, error)
	StartObject() error
	Property(name string) error
	EndObject() error
	// length==-1 => check HaveNext() after every value
	// length==0 => no need to call EndArray
	StartArray() (length int, err error)
	// returns false when array ends or when StartArray returned length!=-1
	HaveNext() (bool, error)
	EndArray() error
	StartMap() (length int, err error)
	NextKey() (string, error)
	EndMap() error
}

type Writer interface {
	SetWriter(writer io.Writer, bufferSize int)
	Int32(value int32) error
	Int64(value int64) error
	Float32(value float32, format byte, precision int) error
	Float64(value float64, format byte, precision int) error
	String(value string) error
	Bytes(value []byte) error
	ByteString(value []byte) error
	StartObject() error
	Property(name string) error
	EndObject() error
	StartArray(length int) error
	EndArray() error
	StartMap(length int) error
	NextKey(key string) error
	EndMap() error
	Flush() error
}

type RawValue interface {
	encoding.TextMarshaler
	gob.GobEncoder
	encoding.TextUnmarshaler
	gob.GobDecoder
}

type InspectorInterface interface {
	LastError() error
	SetReader(io.Reader)
	SetWriter(writer io.Writer, bufferSize int)
	Int32(value *int32)
	Int64(value *int64)
	Int(value *int)
	Float32(value *float32, format byte, precision int)
	Float64(value *float64, format byte, precision int)
	String(value *string)
	Bytes(value *[]byte)
	ByteString(value *[]byte)
	Value(value RawValue)
	StartObject(name string, description string)
	Property(name string, mandatory bool, description string) bool
	PropertyInt32(name string, value *int32, mandatory bool, description string)
	PropertyInt64(name string, value *int64, mandatory bool, description string)
	PropertyInt(name string, value *int, mandatory bool, description string)
	PropertyFloat32(name string, value *float32, format byte, precision int, mandatory bool, description string)
	PropertyFloat64(name string, value *float64, format byte, precision int, mandatory bool, description string)
	PropertyString(name string, value *string, mandatory bool, description string)
	PropertyBytes(name string, value *[]byte, mandatory bool, description string)
	PropertyByteString(name string, value *[]byte, mandatory bool, description string)
	PropertyValue(name string, value RawValue, mandatory bool, description string)
	EndObject()
	ReadArray() int
	WriteArray(name string, elementName string, length int, description string)
	HaveNext() bool
	EndArray()
	ReadMap() int
	WriteMap(name string, elementName string, length int, description string)
	ReadNextKey() string
	WriteNextKey(key string)
	EndMap()
	IsReading() bool
	Flush()
}

type Inspectable interface {
	Inspect(*Inspector)
}

type InspectablePtr[T any] interface {
	Inspectable
	*T
}
