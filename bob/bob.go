// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {

	isYelling := false
	isQuestion := false
	isEmpty := true

	if remark == strings.ToUpper(remark) {
		letterFound := false
		for _, char := range remark {
			if unicode.IsLetter(char) {
				letterFound = true
			}
		}
		isYelling = letterFound
		isEmpty = !letterFound
	}

	if remark[len(remark)-1:] == "?" {
		isQuestion = true
	}

	if isYelling && isQuestion {
		return "Calm down, I know what I'm doing!"
	} else if isEmpty {
		return "Fine. Be that way!"
	} else if isYelling {
		return "Whoa, chill out!"
	} else if isQuestion {
		return "Sure."
	} else {
		return "Whatever."
	}
}
