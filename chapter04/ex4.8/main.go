// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
/*
Exercise 4.8: Modify charcount to count letters, digits, and so on in their Unicode categories,
using functions like unicode.IsLetter.
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[string]int)  // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsControl(r) {
			counts["Control"]++
			utflen[n]++
		}

		if unicode.IsDigit(r) {
			counts["Digit"]++
			utflen[n]++
		}
		if unicode.IsGraphic(r) {
			counts["Graphic"]++
			utflen[n]++
		}
		if unicode.IsLetter(r) {
			counts["Letter"]++
			utflen[n]++
		}
		if unicode.IsLower(r) {
			counts["Lower"]++
			utflen[n]++
		}
		if unicode.IsMark(r) {
			counts["Mark"]++
			utflen[n]++
		}
		if unicode.IsNumber(r) {
			counts["Number"]++
			utflen[n]++

		}
		if unicode.IsPrint(r) {
			counts["Print"]++
			utflen[n]++

		}
		if unicode.IsPunct(r) {
			counts["Punct"]++
			utflen[n]++

		}
		if unicode.IsSpace(r) {
			counts["Space"]++
			utflen[n]++

		}
		if unicode.IsSymbol(r) {
			counts["Symbol"]++
			utflen[n]++

		}
		if unicode.IsTitle(r) {
			counts["Title"]++
			utflen[n]++

		}
		if unicode.IsUpper(r) {
			counts["Upper"]++
			utflen[n]++

		}
	}
	fmt.Printf("kind\tcount\n")
	for c, n := range counts {
		fmt.Printf("%s\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//!-
