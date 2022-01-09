package inspect

import "io"

type Inspector struct {
	impl InspectorInterface
}

func NewInspector(impl InspectorInterface) *Inspector {
	return &Inspector{impl}
}
func (i *Inspector) LastError() error {
	return i.impl.LastError()
}
func (i *Inspector) SetReader(reader io.Reader) {
	i.impl.SetReader(reader)
}
func (i *Inspector) SetWriter(writer io.Writer, bufferSize int) {
	i.impl.SetWriter(writer, bufferSize)
}
func (i *Inspector) Int32(value *int32) {
	i.impl.Int32(value)
}
func (i *Inspector) Int64(value *int64) {
	i.impl.Int64(value)
}
func (i *Inspector) Int(value *int) {
	i.impl.Int(value)
}
func (i *Inspector) Float32(value *float32, format byte, precision int) {
	i.impl.Float32(value, format, precision)
}
func (i *Inspector) Float64(value *float64, format byte, precision int) {
	i.impl.Float64(value, format, precision)
}
func (i *Inspector) String(value *string) {
	i.impl.String(value)
}
func (i *Inspector) Bytes(value *[]byte) {
	i.impl.Bytes(value)
}
func (i *Inspector) ByteString(value *[]byte) {
	i.impl.ByteString(value)
}
func (i *Inspector) Value(value RawValue) {
	i.impl.Value(value)
}

func (i *Inspector) ReadArray() int {
	return i.impl.ReadArray()
}
func (i *Inspector) WriteArray(name string, elementName string, length int, description string) {
	i.impl.WriteArray(name, elementName, length, description)
}
func (i *Inspector) HaveNext() bool {
	return i.impl.HaveNext()
}
func (i *Inspector) EndArray() {
	i.impl.EndArray()
}
func (i *Inspector) ReadMap() int {
	return i.impl.ReadMap()
}
func (i *Inspector) WriteMap(name string, elementName string, length int, description string) {
	i.impl.WriteMap(name, elementName, length, description)
}
func (i *Inspector) ReadNextKey() string {
	return i.impl.ReadNextKey()
}
func (i *Inspector) WriteNextKey(key string) {
	i.impl.WriteNextKey(key)
}
func (i *Inspector) EndMap() {
	i.impl.EndMap()
}
func (i *Inspector) IsReading() bool {
	return i.impl.IsReading()
}
func (i *Inspector) Flush() {
	i.impl.Flush()
}
func (i *Inspector) StartObject(name string, description string) *ObjectInspector {
	i.impl.StartObject(name, description)
	return (*ObjectInspector)(i)
}

type ObjectInspector Inspector

func (o *ObjectInspector) IsReading() bool {
	return o.impl.IsReading()
}
func (o *ObjectInspector) Flush() {
	o.impl.Flush()
}
func (o *ObjectInspector) Property(name string, mandatory bool, description string) *Inspector {
	if o.impl.Property(name, mandatory, description) {
		return (*Inspector)(o)
	}
	return nil
}
func (o *ObjectInspector) Int32(name string, value *int32, mandatory bool, description string) {
	o.impl.PropertyInt32(name, value, mandatory, description)
}
func (o *ObjectInspector) Int64(name string, value *int64, mandatory bool, description string) {
	o.impl.PropertyInt64(name, value, mandatory, description)
}
func (o *ObjectInspector) Int(name string, value *int, mandatory bool, description string) {
	o.impl.PropertyInt(name, value, mandatory, description)
}
func (o *ObjectInspector) Float32(name string, value *float32, format byte, precision int, mandatory bool, description string) {
	o.impl.PropertyFloat32(name, value, format, precision, mandatory, description)
}
func (o *ObjectInspector) Float64(name string, value *float64, format byte, precision int, mandatory bool, description string) {
	o.impl.PropertyFloat64(name, value, format, precision, mandatory, description)
}
func (o *ObjectInspector) String(name string, value *string, mandatory bool, description string) {
	o.impl.PropertyString(name, value, mandatory, description)
}
func (o *ObjectInspector) Bytes(name string, value *[]byte, mandatory bool, description string) {
	o.impl.PropertyBytes(name, value, mandatory, description)
}
func (o *ObjectInspector) ByteString(name string, value *[]byte, mandatory bool, description string) {
	o.impl.PropertyByteString(name, value, mandatory, description)
}
func (o *ObjectInspector) Value(name string, value RawValue, mandatory bool, description string) {
	o.impl.PropertyValue(name, value, mandatory, description)
}
func (o *ObjectInspector) End() *Inspector {
	o.impl.EndObject()
	return (*Inspector)(o)
}
