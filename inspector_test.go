package inspect_test

import (
	"bytes"
	"testing"

	"strconv"

	"github.com/tvanomr/inspect"
	"github.com/tvanomr/inspect/json"
)

type intValue int

func (v *intValue) Inspect(inspector *inspect.Inspector) {
	inspector.Int((*int)(v))
}

const defaultIntValue intValue = 10

func TestValueWrite(t *testing.T) {
	writer := inspect.NewInspector(new(inspect.TextWriteInspector[json.Writer]))
	value := defaultIntValue
	var buffer bytes.Buffer
	writer.SetWriter(&buffer, 10)
	value.Inspect(writer)
	expected := strconv.Itoa(int(defaultIntValue))
	writer.Flush()
	result := string(buffer.Bytes())
	if result != expected {
		t.Fatal("got", result, "expected", expected)
	} else {
		t.Log("ok", result)
	}
}
