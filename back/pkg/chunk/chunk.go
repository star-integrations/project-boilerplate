package chunk

// Index - index
type Index struct {
	From, To int
}

// Chunks - returns an index of the specified size length
func Chunks(length int, chunkSize int) <-chan Index {
	ch := make(chan Index)

	go func() {
		defer close(ch)

		for i := 0; i < length; i += chunkSize {
			idx := Index{i, i + chunkSize}
			if length < idx.To {
				idx.To = length
			}
			ch <- idx
		}
	}()

	return ch
}
