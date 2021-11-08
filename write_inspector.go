package inspect

import "strconv"

type writeInspector[R Writer] struct {
	writer    R
	lastError error
}

func (w *writeInspector[R]) LastError() error {
	return w.lastError
}
func (w *writeInspector[R]) SetBuffer(buffer []byte) {
	w.writer.SetBuffer(buffer)
	w.lastError = nil
}
func (w *writeInspector[R]) Buffer() []byte {
	return w.writer.Buffer()
}
func (w *writeInspector[R]) Int32(value *int32) {
	if w.lastError == nil {
		w.lastError = w.writer.Int32(*value)
	}
}
func (w *writeInspector[R]) Int64(value *int64) {
	if w.lastError == nil {
		w.lastError = w.writer.Int64(*value)
	}
}
func (w *writeInspector[R]) Int(value *int) {
	if w.lastError != nil {
		return
	}
	if strconv.IntSize == 64 {
		w.lastError = w.writer.Int64(int64(*value))
	} else if strconv.IntSize == 32 {
		w.lastError = w.writer.Int32(int32(*value))
	}
}
func (w *writeInspector[R]) Float32(value *float32) {
	if w.lastError == nil {
		w.lastError = w.writer.Float32(*value)
	}
}
func (w *writeInspector[R]) Float64(value *float64) {
	if w.lastError == nil {
		w.lastError = w.writer.Float64(*value)
	}
}
func (w *writeInspector[R]) String(value *string) {
	if w.lastError == nil {
		w.lastError = w.writer.String(*value)
	}
}
func (w *writeInspector[R]) Bytes(value *[]byte) {
	if w.lastError == nil {
		w.lastError = w.writer.Bytes(*value)
	}
}
func (w *writeInspector[R]) StartObject(name string, description string) {
	if w.lastError == nil {
		w.lastError = w.writer.StartObject()
	}
}
func (w *writeInspector[R]) Property(name string, mandatory bool, description string) bool {
	if w.lastError != nil {
		return false
	}
	w.lastError = w.writer.Property(name)
	return true
}
func (w *writeInspector[R]) PropertyInt32(name string, value *int32, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = w.writer.Property(name)
	w.Int32(value)
}
func (w *writeInspector[R]) PropertyInt64(name string, value *int64, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = w.writer.Property(name)
	w.Int64(value)
}
func (w *writeInspector[R]) PropertyInt(name string, value *int, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = w.writer.Property(name)
	w.Int(value)
}
func (w *writeInspector[R]) PropertyFloat32(name string, value *float32, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = w.writer.Property(name)
	w.Float32(value)
}
func (w *writeInspector[R]) PropertyFloat64(name string, value *float64, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = w.writer.Property(name)
	w.Float64(value)
}
func (w *writeInspector[R]) PropertyString(name string, value *string, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = w.writer.Property(name)
	w.String(value)
}
func (w *writeInspector[R]) PropertyBytes(name string, value *[]byte, mandatory bool, description string) {
	if w.lastError != nil {
		return
	}
	w.lastError = w.writer.Property(name)
	w.Bytes(value)
}
func (w *writeInspector[R]) EndObject() {
	if w.lastError == nil {
		w.lastError = w.writer.EndObject()
	}
}
func (w *writeInspector[R]) ReadArray(name string, elementName string, description string) int {
	if w.lastError == nil {
		w.lastError = ErrWriterCantRead
	}
	return 0
}
func (w *writeInspector[R]) WriteArray(name string, elementName string, length int, description string) {
	if w.lastError == nil {
		w.lastError = w.writer.StartArray(length)
	}
}
func (w *writeInspector[R]) EndArray() {
	if w.lastError == nil {
		w.lastError = w.writer.EndArray()
	}
}
func (w *writeInspector[R]) ReadMap(name string, elementName string) int {
	if w.lastError == nil {
		w.lastError = ErrWriterCantRead
	}
	return 0
}
func (w *writeInspector[R]) WriteMap(name string, elementName string, length int, description string) {
	if w.lastError == nil {
		w.lastError = w.writer.StartMap(length)
	}
}
func (w *writeInspector[R]) ReadNextKey() string {
	if w.lastError == nil {
		w.lastError = ErrWriterCantRead
	}
	return ""
}
func (w *writeInspector[R]) WriteNextKey(key string) {
	if w.lastError == nil {
		w.lastError = w.writer.NextKey(key)
	}
}
func (w *writeInspector[R]) EndMap() {
	if w.lastError == nil {
		w.lastError = w.writer.EndMap()
	}
}
func (w *writeInspector[R]) IsReading() bool {
	return false
}
