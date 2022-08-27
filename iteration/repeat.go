package iteration

func Repeat(letter string, counter int) string {
	var repeated string
	for i := 0; i < counter; i++ {
		repeated += letter
	}
	return repeated
}
