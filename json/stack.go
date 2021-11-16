package json

type stack[T any] []T

func (s *stack[T]) push(value T) {
	*s = append(*s, value)
}

func (s *stack[T]) pop() T {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}
