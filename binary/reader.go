package binary

import (
	"bufio"
	"encoding/binary"
	"io"
	"math"
)

type Reader struct {
	reader *bufio.Reader
}

func (r *Reader) SetReader(reader io.Reader) {
	r.reader = bufio.NewReader(reader)
}
func (r *Reader) Int32() (int32, error) {
	result, err := binary.ReadVarint(r.reader)
	if err != nil {
		return 0, err
	}
	return int32(result), nil
}
func (r *Reader) Int64() (int64, error) {
	result, err := binary.ReadVarint(r.reader)
	if err != nil {
		return 0, err
	}
	return result, nil
}
func (r *Reader) Float32() (float32, error) {
	result, err := binary.ReadUvarint(r.reader)
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(uint32(result)), nil
}
func (r *Reader) Float64() (float64, error) {
	result, err := binary.ReadUvarint(r.reader)
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(result), err
}
func (r *Reader) String() (string, error) {
	length, err := binary.ReadVarint(r.reader)
	if err != nil {
		return "", err
	}
	buffer := make([]byte, length)
	read, err := io.ReadFull(r.reader, buffer)
	if err != nil {
		return "", err
	}
	if read < int(length) {
		return "", ErrShortRead
	}
	return string(buffer), nil
}
func (r *Reader) Bytes() ([]byte, error) {
	length, err := binary.ReadVarint(r.reader)
	if err != nil {
		return nil, err
	}
	buffer := make([]byte, length)
	read, err := io.ReadFull(r.reader, buffer)
	if err != nil {
		return nil, err
	}
	if read < int(length) {
		return nil, ErrShortRead
	}
	return buffer, nil
}
func (r *Reader) ByteString() ([]byte, error) {
	return r.Bytes()
}
func (r *Reader) StartObject() error {
	return nil
}
func (r *Reader) Property(name string) error {
	return nil
}
func (r *Reader) EndObject() error {
	return nil
}
func (r *Reader) StartArray() (length int, err error) {
	result, err := binary.ReadVarint(r.reader)
	if err != nil {
		return 0, err
	}
	return int(result), nil
}
func (r *Reader) HaveNext() (bool, error) {
	return false, nil
}
func (r *Reader) EndArray() error {
	return nil
}
func (r *Reader) StartMap() (length int, err error) {
	return r.StartArray()
}
func (r *Reader) NextKey() (string, error) {
	return r.String()
}
func (r *Reader) EndMap() error {
	return nil
}
