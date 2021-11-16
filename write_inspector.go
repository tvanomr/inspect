package inspect

import (
	"io"
	"strconv"
)

type WriteInspector[W Writer] struct {
	writer    W
	lastError error
}

func (w *WriteInspector[W]) LastError() error {
	return w.lastError
}
func (w *WriteInspector[W]) SetWriter(writer io.Writer, bufferSize int) {
	w.writer.SetWriter(writer, bufferSize)
	w.lastError = nil
}
func (w *WriteInspector[W]) SetReader(io.Reader) {
	if w.lastError == nil {
		w.lastError = ErrWriterCantRead
	}
}
func (w *WriteInspector[W]) Int32(value *int32) {
	if w.lastError == nil {
		w.lastError = w.writer.Int32(*value)
	}
}
func (w *WriteInspector[W]) Int64(value *int64) {
	if w.lastError == nil {
		w.lastError = w.writer.Int64(*value)
	}
}
func (w *WriteInspector[W]) Int(value *int) {
	if w.lastError != nil {
		return
	}
	if strconv.IntSize == 64 {
		w.lastError = w.writer.Int64(int64(*value))
	} else if strconv.IntSize == 32 {
		w.lastError = w.writer.Int32(int32(*value))
	}
}
func (w *WriteInspector[W]) Float32(value *float32, format byte, precision int) {
	if w.lastError == nil {
		w.lastError = w.writer.Float32(*value, format, precision)
	}
}
func (w *WriteInspector[W]) Float64(value *float64, format byte, precision int) {
	if w.lastError == nil {
		w.lastError = w.writer.Float64(*value, format, precision)
	}
}
func (w *WriteInspector[W]) String(value *string) {
	if w.lastError == nil {
		w.lastError = w.writer.String(*value)
	}
}
func (w *WriteInspector[W]) Bytes(value *[]byte) {
	if w.lastError == nil {
		w.lastError = w.writer.Bytes(*value)
	}
}
func (w *WriteInspector[W]) ByteString(value *[]byte) {
	if w.lastError == nil {
		w.lastError = w.writer.ByteString(*value)
	}
}
func (w *WriteInspector[W]) StartObject(name string, description string) {
	if w.lastError == nil {
		w.lastError = w.writer.StartObject()
	}
}
func (w *WriteInspector[W]) Property(name string, mandatory bool, description string) bool {
	if w.lastError != nil {
		return false
	}
	w.lastError = w.writer.Property(name)
	return true
}
func (w *WriteInspector[W]) PropertyInt32(name string, value *int32, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = w.writer.Property(name)
	w.Int32(value)
}
func (w *WriteInspector[W]) PropertyInt64(name string, value *int64, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = w.writer.Property(name)
	w.Int64(value)
}
func (w *WriteInspector[W]) PropertyInt(name string, value *int, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = w.writer.Property(name)
	w.Int(value)
}
func (w *WriteInspector[W]) PropertyFloat32(name string, value *float32, format byte, precision int, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = w.writer.Property(name)
	w.Float32(value, format, precision)
}
func (w *WriteInspector[W]) PropertyFloat64(name string, value *float64, format byte, precision int, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = w.writer.Property(name)
	w.Float64(value, format, precision)
}
func (w *WriteInspector[W]) PropertyString(name string, value *string, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = w.writer.Property(name)
	w.String(value)
}
func (w *WriteInspector[W]) PropertyBytes(name string, value *[]byte, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = w.writer.Property(name)
	w.Bytes(value)
}
func (w *WriteInspector[W]) PropertyByteString(name string, value *[]byte, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = w.writer.Property(name)
	w.ByteString(value)
}
func (w *WriteInspector[W]) EndObject() {
	if w.lastError == nil {
		w.lastError = w.writer.EndObject()
	}
}
func (w *WriteInspector[W]) ReadArray(name string, elementName string, description string) int {
	if w.lastError == nil {
		w.lastError = ErrWriterCantRead
	}
	return 0
}
func (w *WriteInspector[W]) WriteArray(name string, elementName string, length int, description string) {
	if w.lastError == nil {
		w.lastError = w.writer.StartArray(length)
	}
}
func (w *WriteInspector[W]) EndArray() {
	if w.lastError == nil {
		w.lastError = w.writer.EndArray()
	}
}
func (w *WriteInspector[W]) ReadMap(name string, elementName string) int {
	if w.lastError == nil {
		w.lastError = ErrWriterCantRead
	}
	return 0
}
func (w *WriteInspector[W]) WriteMap(name string, elementName string, length int, description string) {
	if w.lastError == nil {
		w.lastError = w.writer.StartMap(length)
	}
}
func (w *WriteInspector[W]) ReadNextKey() string {
	if w.lastError == nil {
		w.lastError = ErrWriterCantRead
	}
	return ""
}
func (w *WriteInspector[W]) WriteNextKey(key string) {
	if w.lastError == nil {
		w.lastError = w.writer.NextKey(key)
	}
}
func (w *WriteInspector[W]) EndMap() {
	if w.lastError == nil {
		w.lastError = w.writer.EndMap()
	}
}
func (w *WriteInspector[W]) IsReading() bool {
	return false
}

func (w *WriteInspector[W]) Flush() {
	if w.lastError == nil {
		w.lastError = w.writer.Flush()
	}
}

type TextWriteInspector[W Writer] struct {
	WriteInspector[W]
}

func (w *TextWriteInspector[W]) Value(value RawValue) {
	if w.lastError != nil {
		return
	}
	var data []byte
	data, w.lastError = value.MarshalText()
	if w.lastError == nil {
		w.lastError = w.writer.ByteString(data)
	}
}

func (w *TextWriteInspector[W]) PropertyValue(name string, value RawValue, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = w.writer.Property(name)
	if w.lastError != nil {
		return
	}
	var data []byte
	data, w.lastError = value.MarshalText()
	if w.lastError == nil {
		w.lastError = w.writer.ByteString(data)
	}
}

type BinaryWriteInspector[W Writer] struct {
	WriteInspector[W]
}

func (w *BinaryWriteInspector[W]) Value(value RawValue) {
	if w.lastError != nil {
		return
	}
	var data []byte
	data, w.lastError = value.GobEncode()
	if w.lastError == nil {
		w.lastError = w.writer.Bytes(data)
	}
}

func (w *BinaryWriteInspector[W]) PropertyValue(name string, value RawValue, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = w.writer.Property(name)
	if w.lastError != nil {
		return
	}
	var data []byte
	data, w.lastError = value.GobEncode()
	if w.lastError == nil {
		w.lastError = w.writer.Bytes(data)
	}
}
