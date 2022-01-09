package inspect

import (
	"io"
	"strconv"
)

type WriterPtr[T any] interface {
	*T
	Writer
}

type WriteInspector[W any, PW WriterPtr[W]] struct {
	writer    W
	lastError error
}

func (w *WriteInspector[W, PW]) LastError() error {
	return w.lastError
}
func (w *WriteInspector[W, PW]) SetWriter(writer io.Writer, bufferSize int) {
	PW(&w.writer).SetWriter(writer, bufferSize)
	w.lastError = nil
}
func (w *WriteInspector[W, PW]) SetReader(io.Reader) {
	if w.lastError == nil {
		w.lastError = ErrWriterCantRead
	}
}
func (w *WriteInspector[W, PW]) Int32(value *int32) {
	if w.lastError == nil {
		w.lastError = PW(&w.writer).Int32(*value)
	}
}
func (w *WriteInspector[W, PW]) Int64(value *int64) {
	if w.lastError == nil {
		w.lastError = PW(&w.writer).Int64(*value)
	}
}
func (w *WriteInspector[W, PW]) Int(value *int) {
	if w.lastError != nil {
		return
	}
	if strconv.IntSize == 64 {
		w.lastError = PW(&w.writer).Int64(int64(*value))
	} else if strconv.IntSize == 32 {
		w.lastError = PW(&w.writer).Int32(int32(*value))
	}
}
func (w *WriteInspector[W, PW]) Float32(value *float32, format byte, precision int) {
	if w.lastError == nil {
		w.lastError = PW(&w.writer).Float32(*value, format, precision)
	}
}
func (w *WriteInspector[W, PW]) Float64(value *float64, format byte, precision int) {
	if w.lastError == nil {
		w.lastError = PW(&w.writer).Float64(*value, format, precision)
	}
}
func (w *WriteInspector[W, PW]) String(value *string) {
	if w.lastError == nil {
		w.lastError = PW(&w.writer).String(*value)
	}
}
func (w *WriteInspector[W, PW]) Bytes(value *[]byte) {
	if w.lastError == nil {
		w.lastError = PW(&w.writer).Bytes(*value)
	}
}
func (w *WriteInspector[W, PW]) ByteString(value *[]byte) {
	if w.lastError == nil {
		w.lastError = PW(&w.writer).ByteString(*value)
	}
}
func (w *WriteInspector[W, PW]) StartObject(name string, description string) {
	if w.lastError == nil {
		w.lastError = PW(&w.writer).StartObject()
	}
}
func (w *WriteInspector[W, PW]) Property(name string, mandatory bool, description string) bool {
	if w.lastError != nil {
		return false
	}
	w.lastError = PW(&w.writer).Property(name)
	return true
}
func (w *WriteInspector[W, PW]) PropertyInt32(name string, value *int32, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = PW(&w.writer).Property(name)
	w.Int32(value)
}
func (w *WriteInspector[W, PW]) PropertyInt64(name string, value *int64, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = PW(&w.writer).Property(name)
	w.Int64(value)
}
func (w *WriteInspector[W, PW]) PropertyInt(name string, value *int, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = PW(&w.writer).Property(name)
	w.Int(value)
}
func (w *WriteInspector[W, PW]) PropertyFloat32(name string, value *float32, format byte, precision int, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = PW(&w.writer).Property(name)
	w.Float32(value, format, precision)
}
func (w *WriteInspector[W, PW]) PropertyFloat64(name string, value *float64, format byte, precision int, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = PW(&w.writer).Property(name)
	w.Float64(value, format, precision)
}
func (w *WriteInspector[W, PW]) PropertyString(name string, value *string, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = PW(&w.writer).Property(name)
	w.String(value)
}
func (w *WriteInspector[W, PW]) PropertyBytes(name string, value *[]byte, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = PW(&w.writer).Property(name)
	w.Bytes(value)
}
func (w *WriteInspector[W, PW]) PropertyByteString(name string, value *[]byte, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = PW(&w.writer).Property(name)
	w.ByteString(value)
}
func (w *WriteInspector[W, PW]) EndObject() {
	if w.lastError == nil {
		w.lastError = PW(&w.writer).EndObject()
	}
}
func (w *WriteInspector[W, PW]) ReadArray() int {
	if w.lastError == nil {
		w.lastError = ErrWriterCantRead
	}
	return 0
}
func (w *WriteInspector[W, PW]) WriteArray(name string, elementName string, length int, description string) {
	if w.lastError == nil {
		w.lastError = PW(&w.writer).StartArray(length)
	}
}
func (w *WriteInspector[W, PW]) EndArray() {
	if w.lastError == nil {
		w.lastError = PW(&w.writer).EndArray()
	}
}
func (w *WriteInspector[W, PW]) ReadMap() int {
	if w.lastError == nil {
		w.lastError = ErrWriterCantRead
	}
	return 0
}
func (w *WriteInspector[W, PW]) WriteMap(name string, elementName string, length int, description string) {
	if w.lastError == nil {
		w.lastError = PW(&w.writer).StartMap(length)
	}
}
func (w *WriteInspector[W, PW]) ReadNextKey() string {
	if w.lastError == nil {
		w.lastError = ErrWriterCantRead
	}
	return ""
}
func (w *WriteInspector[W, PW]) WriteNextKey(key string) {
	if w.lastError == nil {
		w.lastError = PW(&w.writer).NextKey(key)
	}
}
func (w *WriteInspector[W, PW]) HaveNext() bool {
	if w.lastError == nil {
		w.lastError = ErrWriterCantRead
	}
	return false
}
func (w *WriteInspector[W, PW]) EndMap() {
	if w.lastError == nil {
		w.lastError = PW(&w.writer).EndMap()
	}
}
func (w *WriteInspector[W, PW]) IsReading() bool {
	return false
}

func (w *WriteInspector[W, PW]) Flush() {
	if w.lastError == nil {
		w.lastError = PW(&w.writer).Flush()
	}
}

type TextWriteInspector[W any, PW WriterPtr[W]] struct {
	WriteInspector[W, PW]
}

func (w *TextWriteInspector[W, PW]) Value(value RawValue) {
	if w.lastError != nil {
		return
	}
	var data []byte
	data, w.lastError = value.MarshalText()
	if w.lastError == nil {
		w.lastError = PW(&w.writer).ByteString(data)
	}
}

func (w *TextWriteInspector[W, PW]) PropertyValue(name string, value RawValue, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = PW(&w.writer).Property(name)
	if w.lastError != nil {
		return
	}
	var data []byte
	data, w.lastError = value.MarshalText()
	if w.lastError == nil {
		w.lastError = PW(&w.writer).ByteString(data)
	}
}

type BinaryWriteInspector[W any, PW WriterPtr[W]] struct {
	WriteInspector[W, PW]
}

func (w *BinaryWriteInspector[W, PW]) Value(value RawValue) {
	if w.lastError != nil {
		return
	}
	var data []byte
	data, w.lastError = value.GobEncode()
	if w.lastError == nil {
		w.lastError = PW(&w.writer).Bytes(data)
	}
}

func (w *BinaryWriteInspector[W, PW]) PropertyValue(name string, value RawValue, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = PW(&w.writer).Property(name)
	if w.lastError != nil {
		return
	}
	var data []byte
	data, w.lastError = value.GobEncode()
	if w.lastError == nil {
		w.lastError = PW(&w.writer).Bytes(data)
	}
}
