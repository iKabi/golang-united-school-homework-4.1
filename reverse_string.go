package reverse_string

func ReverseString(input string) (output string) {
	r := []rune(input)
	for f, b := 0, len(r) - 1; f < b; f, b = f+1, b-1 {
		r[f], r[b] = r[b], r[f]
	}
	return string(r)
}
