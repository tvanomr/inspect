package inspect

import "strconv"

type readInspector[R Reader] struct {
	reader    R
	lastError error
}

func (r *readInspector[R]) LastError() error {
	return r.lastError
}

func (r *readInspector[R]) SetBuffer(buffer []byte) {
	r.reader.SetBuffer(buffer)
	r.lastError = nil
}

func (r *readInspector[R]) Buffer() []byte {
	return r.reader.Buffer()
}

func (r *readInspector[R]) Int32(value *int32) {
	if r.lastError != nil {
		return
	}
	var result int32
	result, r.lastError = r.reader.Int32()
	if r.lastError == nil {
		*value = result
	}
}

func (r *readInspector[R]) Int64(value *int64) {
	if r.lastError != nil {
		return
	}
	var result int64
	result, r.lastError = r.reader.Int64()
	if r.lastError == nil {
		*value = result
	}
}

func (r *readInspector[R]) Int(value *int) {
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

func (r *readInspector[R]) Float32(value *float32) {
	if r.lastError != nil {
		return
	}
	var result float32
	result, r.lastError = r.reader.Float32()
	if r.lastError == nil {
		*value = result
	}
}

func (r *readInspector[R]) Float64(value *float64) {
	if r.lastError != nil {
		return
	}
	var result float64
	result, r.lastError = r.reader.Float64()
	if r.lastError == nil {
		*value = result
	}
}

func (r *readInspector[R]) String(value *string) {
	if r.lastError != nil {
		return
	}
	var result string
	result, r.lastError = r.reader.String()
	if r.lastError == nil {
		*value = result
	}
}

func (r *readInspector[R]) Bytes(value *[]byte) {
	if r.lastError != nil {
		return
	}
	var result []byte
	result, r.lastError = r.reader.Bytes()
	if r.lastError == nil {
		*value = result
	}
}

func (r *readInspector[R]) StartObject(name string, description string) {
	if r.lastError != nil {
		return
	}
	r.lastError = r.reader.StartObject()
}

func (r *readInspector[R]) Property(name string, mandatory bool, description string) bool {
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

func (r *readInspector[R]) PropertyInt32(name string, value *int32, mandatory bool, description string) {
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

func (r *readInspector[R]) PropertyInt64(name string, value *int64, mandatory bool, description string) {
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

func (r *readInspector[R]) PropertyInt(name string, value *int, mandatory bool, description string) {
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

func (r *readInspector[R]) PropertyFloat32(name string, value *float32, mandatory bool, description string) {
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

func (r *readInspector[R]) PropertyFloat64(name string, value *float64, mandatory bool, description string) {
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

func (r *readInspector[R]) PropertyString(name string, value *string, mandatory bool, description string) {
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

func (r *readInspector[R]) PropertyBytes(name string, value *[]byte, mandatory bool, description string) {
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

func (r *readInspector[R]) EndObject() {
	if r.lastError != nil {
		return
	}
	r.lastError = r.reader.EndObject()
}

func (r *readInspector[R]) ReadArray(name string, elementName string, description string) int {
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

func (r *readInspector[R]) WriteArray(name string, elementName string, length int, description string) {
	if r.lastError == nil {
		r.lastError = ErrReaderCantWrite
	}
}

func (r *readInspector[R]) EndArray() {
	if r.lastError == nil {
		r.lastError = r.reader.EndArray()
	}
}

func (r *readInspector[R]) ReadMap(name string, elementName string, description string) int {
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

func (r *readInspector[R]) WriteMap(name string, elementName string, length int, description string) {
	if r.lastError == nil {
		r.lastError = ErrReaderCantWrite
	}
}

func (r *readInspector[R]) ReadNextKey() string {
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

func (r *readInspector[R]) WriteNextKey(key string) {
	if r.lastError == nil {
		r.lastError = ErrReaderCantWrite
	}
}

func (r *readInspector[R]) EndMap() {
	if r.lastError == nil {
		r.lastError = r.reader.EndMap()
	}
}

func (r *readInspector[R]) IsReading() bool {
	return true
}
