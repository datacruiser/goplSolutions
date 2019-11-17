/*
Modify the echo program to print the index and value of each of its arguments,
one per line.
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s string
	for index, arg := range os.Args[1:] {
		s = strconv.Itoa(index) + ", " + arg
		fmt.Println(s)
	}
}
