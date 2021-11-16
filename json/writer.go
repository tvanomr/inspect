package json

import (
	"encoding/base64"
	"io"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/tvanomr/inspect"
)

type Writer struct {
	stream     *jsoniter.Stream
	started    bool
	bufferSize int
}

func (w *Writer) SetWriter(writer io.Writer, bufferSize int) {
	if w.stream == nil {
		w.stream = jsoniter.NewStream(jsoniter.ConfigCompatibleWithStandardLibrary, writer, 10)
	} else {
		w.stream.Reset(writer)
		w.bufferSize = bufferSize
	}
}

func (w *Writer) addComma() {
	if w.started {
		w.stream.WriteMore()
	} else {
		w.started = true
	}
}

func (w *Writer) Int32(value int32) error {
	w.addComma()
	w.stream.WriteInt32(value)
	return nil
}
func (w *Writer) Int64(value int64) error {
	w.addComma()
	w.stream.WriteInt64(value)
	return nil
}
func (w *Writer) Float32(value float32, format byte, precision int) error {
	w.addComma()
	w.stream.SetBuffer(strconv.AppendFloat(w.stream.Buffer(), float64(value), format, precision, 32))
	return nil
}
func (w *Writer) Float64(value float64, format byte, precision int) error {
	w.addComma()
	w.stream.SetBuffer(strconv.AppendFloat(w.stream.Buffer(), value, format, precision, 64))
	return nil
}
func (w *Writer) String(value string) error {
	w.addComma()
	w.stream.WriteString(value)
	return nil
}
func (w *Writer) Bytes(value []byte) error {
	w.addComma()
	w.stream.WriteString(base64.RawURLEncoding.EncodeToString(value))
	return nil
}
func (w *Writer) ByteString(value []byte) error {
	w.addComma()
	w.stream.WriteString(string(value))
	return nil
}
func (w *Writer) StartObject() error {
	w.addComma()
	w.stream.WriteObjectStart()
	w.started = false
	return nil
}
func (w *Writer) Property(name string) error {
	w.addComma()
	w.stream.WriteObjectField(name)
	return nil
}
func (w *Writer) EndObject() error {
	w.stream.WriteObjectEnd()
	w.started = true
	if w.stream.Buffered() > w.bufferSize {
		return w.stream.Flush()
	}
	return nil
}
func (w *Writer) StartArray(length int) error {
	w.addComma()
	w.stream.WriteArrayStart()
	w.started = false
	return nil
}
func (w *Writer) EndArray() error {
	w.stream.WriteArrayEnd()
	w.started = true
	if w.stream.Buffered() > w.bufferSize {
		return w.stream.Flush()
	}
	return nil
}
func (w *Writer) StartMap(length int) error {
	w.addComma()
	w.stream.WriteObjectStart()
	w.started = false
	return nil
}
func (w *Writer) NextKey(key string) error {
	w.addComma()
	w.stream.WriteObjectField(key)
	return nil
}
func (w *Writer) EndMap() error {
	w.stream.WriteObjectEnd()
	w.started = true
	if w.stream.Buffered() > w.bufferSize {
		return w.stream.Flush()
	}
	return nil
}

func (w *Writer) Flush() error {
	return w.stream.Flush()
}

func init() {
	var _ inspect.Writer = (*Writer)(nil)
}
