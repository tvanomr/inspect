package binary

import (
	"encoding/binary"
	"io"
	"math"
)

type Writer struct {
	writer        io.Writer
	variantBuffer []byte
}

func (w *Writer) SetWriter(writer io.Writer, bufferSize int) {
	w.writer = writer
	w.variantBuffer = make([]byte, 10)
}

func (w *Writer) writeBuffer(length int) error {
	written, err := w.writer.Write(w.variantBuffer[:length])
	if err != nil {
		return err
	}
	if written < length {
		return ErrShortWrite
	}
	return nil
}

func (w *Writer) Int32(value int32) error {
	return w.writeBuffer(binary.PutVarint(w.variantBuffer, int64(value)))
}
func (w *Writer) Int64(value int64) error {
	return w.writeBuffer(binary.PutVarint(w.variantBuffer, value))
}
func (w *Writer) Float32(value float32, format byte, precision int) error {
	return w.writeBuffer(binary.PutUvarint(w.variantBuffer, uint64(math.Float32bits(value))))
}
func (w *Writer) Float64(value float64, format byte, precision int) error {
	return w.writeBuffer(binary.PutUvarint(w.variantBuffer, math.Float64bits(value)))
}
func (w *Writer) String(value string) error {
	length := len(value)
	err := w.Int64(int64(length))
	if err != nil {
		return err
	}
	written, err := w.writer.Write([]byte(value))
	if err != nil {
		return err
	}
	if written < length {
		return ErrShortWrite
	}
	return nil
}
func (w *Writer) Bytes(value []byte) error {
	length := len(value)
	err := w.Int64(int64(length))
	if err != nil {
		return err
	}
	written, err := w.writer.Write(value)
	if err != nil {
		return err
	}
	if written < length {
		return ErrShortWrite
	}
	return nil
}
func (w *Writer) ByteString(value []byte) error {
	return w.Bytes(value)
}
func (w *Writer) StartObject() error {
	return nil
}
func (w *Writer) Property(name string) error {
	return nil
}
func (w *Writer) EndObject() error {
	return nil
}
func (w *Writer) StartArray(length int) error {
	return w.Int64(int64(length))
}
func (w *Writer) EndArray() error {
	return nil
}
func (w *Writer) StartMap(length int) error {
	return w.Int64(int64(length))
}
func (w *Writer) NextKey(key string) error {
	return w.String(key)
}
func (w *Writer) EndMap() error {
	return nil
}
func (w *Writer) Flush() error {
	//this writer is not buffered
	return nil
}
