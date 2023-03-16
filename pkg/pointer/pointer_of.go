package pointer

func Of[T any](something T) *T {
	return &something
}
