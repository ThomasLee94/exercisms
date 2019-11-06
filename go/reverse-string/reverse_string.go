package reverse

// reverse the order of input string
func Reverse(str string) string {
	// case: string is empty
	if str == "" {
		return ""
	}

	reverse_str := ""

	for _, char := range str {
		reverse_str = string(char) + reverse_str
	}

	return string(reverse_str)

}
