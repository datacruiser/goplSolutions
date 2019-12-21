/*
Exercise 4.7: Modify reverse to reverse the characters of
a []byte slice that represents a UTF-8-encoded string, in place. Can you do it without allocating new memory?
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseUTF8(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverse(b[i : i+size])
		i += size
	}
	reverseASCII(b)
	return b
}

func reverseASCII(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}

func reverse(s []byte) {
	var p, q rune        // first and last runes
	var psize, qsize int // sizes of first and last runes

	for i, j := 0, len(s); i < j-1; i, j = i+qsize, j-psize {
		p, psize = utf8.DecodeRune(s[i:])
		q, qsize = utf8.DecodeLastRune(s[:j])

		// shift bytes between first and last runes
		//if size of last > first

		if qsize > psize {
			copy(s[i+qsize:], s[i+psize:j-qsize])
		}

		//copy first rune to first
		copy(s[i:], []byte(string(q)))

		// copy last rune to first to last
		copy(s[j-psize:], []byte(string(p)))
	}
}

func main() {
	b := []byte("Hello, 世界")
	s := []byte("Hello, world!")
	fmt.Println(string(b))
	reverse(b)
	fmt.Println(string(b))
	fmt.Println(string(reverseUTF8(b)))
	fmt.Println(string(reverseASCII(b)))
	fmt.Println(string(reverseASCII(s)))
}
