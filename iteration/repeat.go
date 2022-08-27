package iteration

const repeatCounter int= 5

func Repeat(letter string) string {
	var repeated string
	for i := 0; i < repeatCounter; i++ {
		repeated += letter
	}
	return repeated
}