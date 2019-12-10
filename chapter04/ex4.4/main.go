/*
Exercise 4.4: Write a version of rotate that operates in a single pass.
*/

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

package main

import "fmt"

func main() {
	//!+array

	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s)
	rotate(s, 3)
	fmt.Println(s)

}

func rotate(s []int, n int) {
	num := n % len(s)
	double := append(s, s[:num]...)
	copy(s, double[num:num+len(s)])
}
