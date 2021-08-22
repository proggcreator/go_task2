package main

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func repiterstr(r rune, n int) string {
	var repitB strings.Builder
	for i := 0; i < n-1; i++ {
		repitB.WriteRune(r)
	}
	return repitB.String()
}

func Unpack(str string) string {
	var previousRune rune
	var b strings.Builder
	escape := '\\'
	flagTwoEscape := false
	flagEscapeDigit := false
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		str = str[size:]
		switch {

		// catching two escape
		case previousRune == escape && r == escape:
			previousRune = r
			flagTwoEscape = true
			b.WriteRune(r)

			//  qwe\\5 must have - qwe\\\\\
		case unicode.IsDigit(r) && flagTwoEscape:
			n, _ := strconv.Atoi(string(r))
			b.WriteString(repiterstr(previousRune, n))
			flagTwoEscape = false

		//   qwe\45 must have - qwe44444
		case unicode.IsDigit(previousRune) && unicode.IsDigit(r) && flagEscapeDigit:
			n, _ := strconv.Atoi(string(r))
			b.WriteString(repiterstr(previousRune, n))
			flagEscapeDigit = false
			previousRune = r

		//  qwe\4\5 must have - qwe45
		case previousRune == escape && unicode.IsDigit(r):
			flagEscapeDigit = true
			b.WriteRune(r)
			previousRune = r

		// a4
		case unicode.IsDigit(r) && unicode.IsLetter(previousRune):
			n, _ := strconv.Atoi(string(r))
			b.WriteString(repiterstr(previousRune, n))
			previousRune = r

		// if letter
		case unicode.IsLetter(r):
			b.WriteRune(r)
			previousRune = r

		// if escape
		case r == escape:
			previousRune = r
		}
	}
	return b.String()
}
