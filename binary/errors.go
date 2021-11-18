package binary

type binError int

const (
	ErrShortRead binError = iota
	ErrShortWrite
)

var binErrorMessages = map[binError]string{
	ErrShortRead:  "short read",
	ErrShortWrite: "short write"}

func (b binError) Error() string {
	return binErrorMessages[b]
}
