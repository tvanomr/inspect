package inspect

type Reader interface {
	SetBuffer(buffer []byte)
	Buffer() []byte
	Int32() (int32, error)
	Int64() (int64, error)
	Float32() (float32, error)
	Float64() (float64, error)
	String() (string, error)
	Bytes() ([]byte, error)
	StartObject() error
	Property(name string) error
	EndObject() error
	StartArray() (length int, err error)
	EndArray() error
	StartMap() (length int, err error)
	NextKey() (string, error)
	EndMap() error
}

type Writer interface {
	Buffer() []byte
	SetBuffer(buffer []byte)
	Int32(value int32) error
	Int64(value int64) error
	Float32(value float32) error
	Float64(value float64) error
	String(value string) error
	Bytes(value []byte) error
	StartObject() error
	Property(name string) error
	EndObject() error
	StartArray(length int) error
	EndArray() error
	StartMap(length int) error
	NextKey(key string) error
	EndMap() error
}

type InspectorInterface interface {
	LastError() error
	SetBuffer(buffer []byte)
	Buffer() []byte
	Int32(value *int32)
	Int64(value *int64)
	Int(value *int)
	Float32(value *float32)
	Float64(value *float64)
	String(value *string)
	Bytes(value *[]byte)
	StartObject(name string, description string)
	Property(name string, mandatory bool, description string) bool
	PropertyInt32(name string, value *int32, mandatory bool, description string)
	PropertyInt64(name string, value *int64, mandatory bool, description string)
	PropertyInt(name string, value *int, mandatory bool, description string)
	PropertyFloat32(name string, value *float32, mandatory bool, description string)
	PropertyFloat64(name string, value *float64, mandatory bool, description string)
	PropertyString(name string, value *string, mandatory bool, description string)
	PropertyBytes(name string, value *[]byte, mandatory bool, description string)
	EndObject()
	ReadArray(name string, elementName string, description string) int
	WriteArray(name string, elementName string, length int, description string)
	EndArray()
	ReadMap(name string, elementName string) int
	WriteMap(name string, elementName string, length int, description string)
	ReadNextKey() string
	WriteNextKey(key string)
	EndMap()
	IsReading() bool
}
