package inspect

import (
	"io"
	"strconv"
)

type ReadInspector[R Reader] struct {
	reader    R
	lastError error
}

func (r *ReadInspector[R]) LastError() error {
	return r.lastError
}

func (r *ReadInspector[R]) SetWriter(io.Writer) {
	if r.lastError == nil {
		r.lastError = ErrReaderCantWrite
	}
}

func (r *ReadInspector[R]) SetReader(reader io.Reader) {
	r.reader.SetReader(reader)
	r.lastError = nil
}

func (r *ReadInspector[R]) Int32(value *int32) {
	if r.lastError != nil {
		return
	}
	var result int32
	result, r.lastError = r.reader.Int32()
	if r.lastError == nil {
		*value = result
	}
}

func (r *ReadInspector[R]) Int64(value *int64) {
	if r.lastError != nil {
		return
	}
	var result int64
	result, r.lastError = r.reader.Int64()
	if r.lastError == nil {
		*value = result
	}
}

func (r *ReadInspector[R]) Int(value *int) {
	if r.lastError != nil {
		return
	}
	if strconv.IntSize == 64 {
		var result int64
		result, r.lastError = r.reader.Int64()
		if r.lastError == nil {
			*value = int(result)
		}
	} else if strconv.IntSize == 32 {
		var result int32
		result, r.lastError = r.reader.Int32()
		if r.lastError == nil {
			*value = int(result)
		}
	}
}

func (r *ReadInspector[R]) Float32(value *float32, format byte, precision int) {
	if r.lastError != nil {
		return
	}
	var result float32
	result, r.lastError = r.reader.Float32()
	if r.lastError == nil {
		*value = result
	}
}

func (r *ReadInspector[R]) Float64(value *float64, format byte, precision int) {
	if r.lastError != nil {
		return
	}
	var result float64
	result, r.lastError = r.reader.Float64()
	if r.lastError == nil {
		*value = result
	}
}

func (r *ReadInspector[R]) String(value *string) {
	if r.lastError != nil {
		return
	}
	var result string
	result, r.lastError = r.reader.String()
	if r.lastError == nil {
		*value = result
	}
}

func (r *ReadInspector[R]) Bytes(value *[]byte) {
	if r.lastError != nil {
		return
	}
	var result []byte
	result, r.lastError = r.reader.Bytes()
	if r.lastError == nil {
		*value = result
	}
}

func (r *ReadInspector[R]) ByteString(value *[]byte) {
	if r.lastError != nil {
		return
	}
	var result []byte
	result, r.lastError = r.reader.ByteString()
	if r.lastError == nil {
		*value = result
	}
}

func (r *ReadInspector[R]) StartObject(name string, description string) {
	if r.lastError != nil {
		return
	}
	r.lastError = r.reader.StartObject()
}

func (r *ReadInspector[R]) Property(name string, mandatory bool, description string) bool {
	if r.lastError != nil {
		return false
	}
	err := r.reader.Property(name)
	if err != nil {
		if mandatory || err != ErrNoField {
			r.lastError = err
		}
	}
	return false
}

func (r *ReadInspector[R]) PropertyInt32(name string, value *int32, mandatory bool, description string) {
	if r.lastError != nil {
		return
	}
	err := r.reader.Property(name)
	if err != nil {
		if mandatory || err != ErrNoField {
			r.lastError = err
		}
		return
	}
	var result int32
	result, r.lastError = r.reader.Int32()
	if r.lastError != nil {
		*value = result
	}
}

func (r *ReadInspector[R]) PropertyInt64(name string, value *int64, mandatory bool, description string) {
	if r.lastError != nil {
		return
	}
	err := r.reader.Property(name)
	if err != nil {
		if mandatory || err != ErrNoField {
			r.lastError = err
		}
		return
	}
	var result int64
	result, r.lastError = r.reader.Int64()
	if r.lastError != nil {
		*value = result
	}
}

func (r *ReadInspector[R]) PropertyInt(name string, value *int, mandatory bool, description string) {
	if r.lastError != nil {
		return
	}
	err := r.reader.Property(name)
	if err != nil {
		if mandatory || err != ErrNoField {
			r.lastError = err
		}
		return
	}
	if strconv.IntSize == 64 {
		var result int64
		result, r.lastError = r.reader.Int64()
		if r.lastError != nil {
			*value = int(result)
		}
	} else if strconv.IntSize == 32 {
		var result int32
		result, r.lastError = r.reader.Int32()
		if r.lastError != nil {
			*value = int(result)
		}
	}
}

func (r *ReadInspector[R]) PropertyFloat32(name string, value *float32, format byte, precision int, mandatory bool, description string) {
	if r.lastError != nil {
		return
	}
	err := r.reader.Property(name)
	if err != nil {
		if mandatory || err != ErrNoField {
			r.lastError = err
		}
		return
	}
	var result float32
	result, r.lastError = r.reader.Float32()
	if r.lastError != nil {
		*value = result
	}
}

func (r *ReadInspector[R]) PropertyFloat64(name string, value *float64, format byte, precision int, mandatory bool, description string) {
	if r.lastError != nil {
		return
	}
	err := r.reader.Property(name)
	if err != nil {
		if mandatory || err != ErrNoField {
			r.lastError = err
		}
		return
	}
	var result float64
	result, r.lastError = r.reader.Float64()
	if r.lastError != nil {
		*value = result
	}
}

func (r *ReadInspector[R]) PropertyString(name string, value *string, mandatory bool, description string) {
	if r.lastError != nil {
		return
	}
	err := r.reader.Property(name)
	if err != nil {
		if mandatory || err != ErrNoField {
			r.lastError = err
		}
		return
	}
	var result string
	result, r.lastError = r.reader.String()
	if r.lastError != nil {
		*value = result
	}
}

func (r *ReadInspector[R]) PropertyBytes(name string, value *[]byte, mandatory bool, description string) {
	if r.lastError != nil {
		return
	}
	err := r.reader.Property(name)
	if err != nil {
		if mandatory || err != ErrNoField {
			r.lastError = err
		}
		return
	}
	var result []byte
	result, r.lastError = r.reader.Bytes()
	if r.lastError != nil {
		*value = result
	}
}

func (r *ReadInspector[R]) PropertyByteString(name string, value *[]byte, mandatory bool, description string) {
	if r.lastError != nil {
		return
	}
	err := r.reader.Property(name)
	if err != nil {
		if mandatory || err != ErrNoField {
			r.lastError = err
		}
		return
	}
	var result []byte
	result, r.lastError = r.reader.ByteString()
	if r.lastError != nil {
		*value = result
	}
}

func (r *ReadInspector[R]) EndObject() {
	if r.lastError != nil {
		return
	}
	r.lastError = r.reader.EndObject()
}

func (r *ReadInspector[R]) ReadArray(name string, elementName string, description string) int {
	if r.lastError != nil {
		return 0
	}
	var result int
	result, r.lastError = r.reader.StartArray()
	if r.lastError != nil {
		return 0
	}
	return result
}

func (r *ReadInspector[R]) WriteArray(name string, elementName string, length int, description string) {
	if r.lastError == nil {
		r.lastError = ErrReaderCantWrite
	}
}

func (r *ReadInspector[R]) HaveNext() bool {
	if r.lastError != nil {
		return false
	}
	var result bool
	result, r.lastError = r.reader.HaveNext()
	if r.lastError != nil {
		return false
	}
	return result
}

func (r *ReadInspector[R]) EndArray() {
	if r.lastError == nil {
		r.lastError = r.reader.EndArray()
	}
}

func (r *ReadInspector[R]) ReadMap(name string, elementName string, description string) int {
	if r.lastError != nil {
		return 0
	}
	var result int
	result, r.lastError = r.reader.StartMap()
	if r.lastError != nil {
		return 0
	}
	return result
}

func (r *ReadInspector[R]) WriteMap(name string, elementName string, length int, description string) {
	if r.lastError == nil {
		r.lastError = ErrReaderCantWrite
	}
}

func (r *ReadInspector[R]) ReadNextKey() string {
	if r.lastError != nil {
		return ""
	}
	var key string
	key, r.lastError = r.reader.NextKey()
	if r.lastError != nil {
		return ""
	}
	return key
}

func (r *ReadInspector[R]) WriteNextKey(key string) {
	if r.lastError == nil {
		r.lastError = ErrReaderCantWrite
	}
}

func (r *ReadInspector[R]) EndMap() {
	if r.lastError == nil {
		r.lastError = r.reader.EndMap()
	}
}

func (r *ReadInspector[R]) IsReading() bool {
	return true
}

type TextReadInspector[R Reader] struct {
	ReadInspector[R]
}

func (r *TextReadInspector[R]) Value(value RawValue) {
	if r.lastError != nil {
		return
	}
	var data []byte
	data, r.lastError = r.reader.ByteString()
	if r.lastError == nil {
		r.lastError = value.UnmarshalText(data)
	}
}

func (r *TextReadInspector[R]) PropertyValue(name string, value RawValue, mandatory bool, description string) {
	if r.lastError != nil {
		return
	}
	err := r.reader.Property(name)
	if err != nil {
		if mandatory || err != ErrNoField {
			r.lastError = err
		}
		return
	}
	var data []byte
	data, r.lastError = r.reader.ByteString()
	if r.lastError != nil {
		r.lastError = value.UnmarshalText(data)
	}
}

type BinaryReadInspector[R Reader] struct {
	ReadInspector[R]
}

func (r *BinaryReadInspector[R]) Value(value RawValue) {
	if r.lastError != nil {
		return
	}
	var data []byte
	data, r.lastError = r.reader.Bytes()
	if r.lastError == nil {
		r.lastError = value.GobDecode(data)
	}
}

func (r *BinaryReadInspector[R]) PropertyValue(name string, value RawValue, mandatory bool, description string) {
	if r.lastError != nil {
		return
	}
	err := r.reader.Property(name)
	if err != nil {
		if mandatory || err != ErrNoField {
			r.lastError = err
		}
		return
	}
	var data []byte
	data, r.lastError = r.reader.Bytes()
	if r.lastError != nil {
		r.lastError = value.GobDecode(data)
	}
}
