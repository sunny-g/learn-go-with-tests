package iteration

// Repeats the given string 5 times.
func Repeat(character string, count int) (repeated string) {
	for i := 0; i < count; i++ {
		repeated += character
	}
	return
}
