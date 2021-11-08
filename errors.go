package inspect

type inspectError int

const (
	ErrNoField inspectError = iota
	ErrReaderCantWrite
	ErrWriterCantRead
)

var errorMessages = map[inspectError]string{
	ErrNoField:         "field not present",
	ErrReaderCantWrite: "trying to write to a reading inspector",
	ErrWriterCantRead:  "trying to read from a a writing inspector"}

func (i inspectError) Error() string {
	return errorMessages[i]
}
