// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

/*
Exercise 3.12: Write a function that reports whether two strings are anagrams of each other,
that is, they contain the same letters in a different order.
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
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "chapter03/ex12: must have 2 arguments.")
		os.Exit(1)
	}
}

//!+

//
func anagram(a, b string) bool {
	return equals(runeOccurrences(a), runeOccurrences(b))
}

func runeOccurrences(s string) map[rune]int {
	occurrences := make(map[rune]int)
	for _, r := range s {
		occurrences[r]++
	}
	return occurrences

}

func equals(m1, m2 map[rune]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}

	return true
}
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
