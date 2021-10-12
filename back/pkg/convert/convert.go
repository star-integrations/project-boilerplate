package convert

// Str2Ptr - string to pointer string
func Str2Ptr(s string) *string {
	return &s
}

// Int2Ptr - int to pointer int
func Int2Ptr(i int) *int {
	return &i
}

// Int642Ptr - int64 to pointer int64
func Int642Ptr(i int64) *int64 {
	return &i
}
