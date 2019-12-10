/*
Exercise 4.1: Write a function that counts the number of bits that are different in two SHA256 hashes.
(See PopCount from Section 2.6.2.)
*/

package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	//fmt.Println(pc)

	if len(os.Args) != 3 {
		fmt.Println(os.Stderr, "ch04/ex01 must have 2 arguments")
		os.Exit(1)
	}

	fmt.Printf("%d\n", sha256PopCount(os.Args[1], os.Args[2]))
}

func sha256PopCount(a, b string) int {
	digestA := sha256.Sum256([]byte(a))
	digestB := sha256.Sum256([]byte(b))
	return popCount(digestA, digestB)

}

func popCount(a, b [32]byte) int {
	pop := 0
	for i := range a {
		pop += int(pc[a[i]^b[i]])
	}
	return pop
}
