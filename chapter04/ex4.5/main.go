/*
Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice.
*/
package main

import "fmt"

func main() {
	s := []string{"A", "A", "B", "C", "C", "D"}
	fmt.Println(removeDup(s))

}

func removeDup(s []string) []string {
	for i := 0; i < len(s)-1; {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
		} else {
			i++
		}
	}
	return s
}
