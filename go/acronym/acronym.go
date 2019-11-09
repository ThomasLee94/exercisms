// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"log"
	"regexp"
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	var outputRune []rune

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	processedStr := reg.ReplaceAllString(s, "")

	words := strings.Split(processedStr, " ")

	for _, word := range words {
		runeSlice := []rune(word)

		if unicode.IsLetter(runeSlice[0]) {
			outputRune = append(outputRune, unicode.ToUpper(runeSlice[0]))
		}

	}

	outputStr := string(outputRune)
	return outputStr
}
