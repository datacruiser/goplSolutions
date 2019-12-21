/*
Exercise 4.6: Write an in-place function that squashes each run of adjacent Unicode spaces (see unicode.IsSpace)
in a UTF-8-encoded []byte slice into a single ASCII space.
*/
package main

import (
	"bytes"
	"fmt"
	"unicode"
	"unicode/utf8"
)

func replaceSpace(b []byte) []byte {
	var buf bytes.Buffer
	for i := 0; i < len(b); {
		r, size := utf8.DecodeRuneInString(string(b[i:]))
		if unicode.IsSpace(r) {
			nextrune, _ := utf8.DecodeRuneInString(string(b[i+size:]))
			if !unicode.IsSpace(nextrune) {
				buf.WriteRune(' ')
			}
		} else {
			buf.WriteRune(r)
		}
		i += size
	}
	return buf.Bytes()
}

func main() {
	s := "1  +  2      =  3"
	fmt.Println(string(replaceSpace([]byte(s))))
}
