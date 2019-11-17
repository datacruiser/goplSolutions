/*
Exercise 1.4: Modify dup2 to print the names of all ﬁles in which each duplicated line occurs.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	countsFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, countsFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, countsFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, countsFiles[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, countsFiles map[string][]string) {
	input := bufio.NewScanner(f)
	name := f.Name()
	for input.Scan() {
		counts[input.Text()]++
		if !fileContains(countsFiles[input.Text()], name) {
			countsFiles[input.Text()] = append(countsFiles[input.Text()], name)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

func fileContains(array []string, value string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}

//!-
