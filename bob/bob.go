// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if remark == "" {
		return "Fine. Be that way!"
	}
	isQuestion := remark[len(remark)-1] == '?'
	hasLetters := false
	isShouting := true
	for _, r := range remark {
		if unicode.IsLetter(r) {
			hasLetters = true
			if !unicode.IsUpper(r) {
				isShouting = false
			}
		}
	}
	if isShouting && hasLetters {
		if isQuestion {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	}
	if isQuestion {
		return "Sure."
	}
	return "Whatever."
}

