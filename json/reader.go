package json

import (
	"encoding/base64"
	"io"

	jsoniter "github.com/json-iterator/go"
)

type Reader struct {
	iterator        *jsoniter.Iterator
	levels          stack[bool]
	isObject        bool
	endReached      bool
	endReachedStack stack[bool]
}

func (r *Reader) SetReader(reader io.Reader) {
	if r.iterator == nil {
		r.iterator = jsoniter.Parse(jsoniter.ConfigCompatibleWithStandardLibrary, reader, 1024)
	} else {
		r.iterator.Reset(reader)
	}
}

func (r *Reader) Int32() (int32, error) {
	return r.iterator.ReadInt32(), r.iterator.Error
}
func (r *Reader) Int64() (int64, error) {
	return r.iterator.ReadInt64(), r.iterator.Error
}
func (r *Reader) Float32() (float32, error) {
	return r.iterator.ReadFloat32(), r.iterator.Error
}

func (r *Reader) Float64() (float64, error) {
	return r.iterator.ReadFloat64(), r.iterator.Error
}

func (r *Reader) String() (string, error) {
	return r.iterator.ReadString(), r.iterator.Error
}

func (r *Reader) Bytes() ([]byte, error) {
	result, err := base64.RawURLEncoding.DecodeString(r.iterator.ReadString())
	if r.iterator.Error != nil {
		return nil, r.iterator.Error
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Reader) ByteString() ([]byte, error) {
	return append([]byte(nil), r.iterator.ReadStringAsSlice()...), r.iterator.Error
}

func (r *Reader) StartObject() error {
	r.levels.push(r.isObject)
	r.isObject = true
	return nil
}
func (r *Reader) Property(name string) error {
	if !r.isObject {
		return ErrNotAnObject
	}
	field := r.iterator.ReadObject()
	if r.iterator.Error != nil {
		return r.iterator.Error
	}
	if len(field) == 0 {
		return ErrUnexpectedObjectEnd
	}
	if field != name {
		return ErrWrongField
	}
	return nil
}
func (r *Reader) EndObject() error {
	if !r.isObject {
		return ErrNotAnObject
	}
	r.isObject = r.levels.pop()
	field := r.iterator.ReadObject()
	if len(field) > 0 {
		return ErrObjectTooBig
	}
	return nil
}
func (r *Reader) StartArray() (length int, err error) {
	arrayHasItems := r.iterator.ReadArray()
	if r.iterator.Error != nil {
		return 0, r.iterator.Error
	}
	if !arrayHasItems {
		return 0, nil
	}
	r.endReachedStack.push(r.endReached)
	r.levels.push(r.isObject)
	r.isObject = false
	r.endReached = false
	return -1, nil
}
func (r *Reader) HaveNext() (bool, error) {
	arrayHasItems := r.iterator.ReadArray()
	if r.iterator.Error != nil {
		return false, r.iterator.Error
	}
	r.endReached = arrayHasItems
	return arrayHasItems, nil
}
func (r *Reader) EndArray() error {
	if !r.endReached {
		r.endReached = r.iterator.ReadArray()
		if !r.endReached {
			r.endReached = r.endReachedStack.pop()
			r.isObject = r.levels.pop()
			return ErrArrayTooBig
		}
	}
	r.endReached = r.endReachedStack.pop()
	r.isObject = r.levels.pop()
	return nil
}
func (r *Reader) StartMap() (length int, err error) {
	r.levels.push(r.isObject)
	r.endReachedStack.push(r.endReached)
	r.isObject = true
	r.endReached = false
	return -1, nil
}
func (r *Reader) NextKey() (string, error) {
	key := r.iterator.ReadObject()
	if r.iterator.Error != nil {
		return "", r.iterator.Error
	}
	if len(key) == 0 {
		r.endReached = true
	}
	return key, nil
}
func (r *Reader) EndMap() error {
	if !r.endReached {
		key := r.iterator.ReadObject()
		if len(key) > 0 {
			r.isObject = r.levels.pop()
			r.endReached = r.endReachedStack.pop()
			return ErrMapTooBig
		}
	}
	r.isObject = r.levels.pop()
	r.endReached = r.endReachedStack.pop()
	return nil
}
