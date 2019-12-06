// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//

/*
Exercise 3.11: Enhance comma so that it deals correctly with ﬂoating-point numbers and an optional sign.
*/
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", commaSigned(os.Args[i]))
	}
}

//!+

//
func commaSigned(s string) string {
	var start, end int

	if strings.HasPrefix(s, "-") {
		start = 1
	} else {
		start = 0
	}

	if strings.Contains(s, ".") {
		end = strings.Index(s, ".")
	} else {
		end = len(s)
	}

	return s[:start] + comma(s[start:end]) + s[end:]
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	i := (3 - utf8.RuneCountInString(s)%3) % 3
	for _, r := range s {
		if i == 3 {
			buf.WriteByte(',')
			i = 0
		}
		buf.WriteRune(r)
		i++
	}

	return buf.String()

}

//!-
