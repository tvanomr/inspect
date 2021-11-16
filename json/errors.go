package json

type jsonError int

const (
	ErrNoWriter jsonError = iota
	ErrUnexpectedArrayEnd
	ErrUnexpectedObjectEnd
	ErrNotAnObject
	ErrWrongField
	ErrObjectTooBig
	ErrArrayTooBig
	ErrMapTooBig
)

var errorMessages = map[jsonError]string{
	ErrNoWriter:            "no io.writer set for a writer",
	ErrUnexpectedArrayEnd:  "unexpected array end",
	ErrUnexpectedObjectEnd: "unexpected object end",
	ErrNotAnObject:         "Property() call outside of object",
	ErrWrongField:          "supplied field name differs from the one in json",
	ErrObjectTooBig:        "object contains more fields than requested",
	ErrArrayTooBig:         "array contains more items than was read",
	ErrMapTooBig:           "map contains more items than was read"}

func (j jsonError) Error() string {
	return errorMessages[j]
}
