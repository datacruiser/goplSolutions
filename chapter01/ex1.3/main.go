/*
Experiment to measure the difference in running time between our potentially
inefﬁcient versions and the one that uses strings.Join.
(Section 1.6 illustrates part of the time package,
and Section 11.4 shows how to write benchmark tests for systematic performance evaluation.)
*/

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1(args []string) {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}

	fmt.Println(s)
	fmt.Printf("%.10fs elapsed\n", time.Since(start).Seconds())

}

func echo2(args []string) {
	s, sep := "", ""
	start := time.Now()
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.10fs elapsed\n", time.Since(start).Seconds())

}

func echo3(args []string) {
	start := time.Now()
	fmt.Println(strings.Join(args[1:], " "))
	fmt.Printf("%.10fs elapsed\n", time.Since(start).Seconds())

}

func main() {
	echo1(os.Args)
	echo2(os.Args)
	echo3(os.Args)
}
