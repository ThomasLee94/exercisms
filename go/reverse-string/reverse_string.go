package reverse

// reverse the order of input string
func Reverse(str string) string {
	// case: string is empty
	if str == "" {
		return ""
	}
	// make array of runes (chars)
	chars := []rune(str)

	// loop through chars array of str's & swap chars
	max_len_index := len(chars)

	for i := 0; i < max_len_index; i += 1 {
		for j := max_len_index; j >= 0; j -= 1 {
			chars[i], chars[j] = chars[j], chars[i]
		}
	}

	return string(chars)

}
