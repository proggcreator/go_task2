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

		//если встретили "//"
		case previousRune == escape && r == escape:
			previousRune = r
			flagTwoEscape = true
			b.WriteRune(r)

			//  qwe\\5 получаем - qwe\\\\\
		case unicode.IsDigit(r) && flagTwoEscape:
			n, _ := strconv.Atoi(string(r))
			b.WriteString(repiterstr(previousRune, n))
			flagTwoEscape = false

		//   qwe\45 получаем - qwe44444
		case unicode.IsDigit(previousRune) && unicode.IsDigit(r) && flagEscapeDigit:
			n, _ := strconv.Atoi(string(r))
			b.WriteString(repiterstr(previousRune, n))
			flagEscapeDigit = false
			previousRune = r

		//  qwe\4\5 получаем - qwe45
		case previousRune == escape && unicode.IsDigit(r):
			flagEscapeDigit = true
			b.WriteRune(r)
			previousRune = r

		// a4 последовательность
		case unicode.IsDigit(r) && unicode.IsLetter(previousRune):
			n, _ := strconv.Atoi(string(r))
			b.WriteString(repiterstr(previousRune, n))
			previousRune = r

		// если буква
		case unicode.IsLetter(r):
			b.WriteRune(r)
			previousRune = r

		// ели последовательность / дальше
		case r == escape:
			previousRune = r
		}
	}
	return b.String()
}
