package isogram

import(
	"errors";
	"unicode";
)

// CHECKS IF PARAMETER STRING IS AN ISOGRAM
func IsogramCheck(word string) (bool, int, error) {
	if word == "" {
		err := errors.New("Please input a string!")
		return false, -1, err
	}

	// CREATE EMPTY MAP DECLARING TYPES 
	dict := make(map[rune]bool)

	// ITERATE OVER INPUT STRING
	for _, char := range word {
		// FORCE EVERYTHING TO LOWERCASE
		char := unicode.ToLower(char)
		// IF THE CHAR IS FOUND IN MAP
		if _, found := dict[char]; found {
			/* 
			IF THE CHAR IS AN ALPHABETICAL CHAR
			*/
		}
	}

}