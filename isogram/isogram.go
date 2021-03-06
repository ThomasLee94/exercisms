package isogram

import (
	"unicode"
)

//TAKES INPUT STRING AND RETURNS A BOOL
func IsIsogram(word string) bool {

	if word == "" {
		return true
	}

	dict := make(map[rune]bool)

	//
	for _, char := range word {
		// index, bool of found
		_, found := dict[unicode.ToUpper(char)]

		if found {
			return false
		}

		if unicode.IsLetter(char) {
			dict[unicode.ToUpper(char)] = false
		}

	}
	return true
}
