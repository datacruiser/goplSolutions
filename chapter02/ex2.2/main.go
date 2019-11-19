// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
/*
Exercise 2.2: Write a general-purpose unit-conversion program analogous to cf
that reads numbers from its command-line arguments or from the standard input
if there are no arguments, and converts each number into units like temperature
in Celsius and Fahrenheit, length in feet and meters, weight in pounds and kilograms, and the like.
*/
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jijiwhywhy/goplSolutions/chapter02/ex2.2/unitconv"
)

func handle(arg string) {
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	{
		k := unitconv.Kelvin(t)
		c := unitconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			k, unitconv.KToC(k), c, unitconv.CToK(c))
	}

	{
		f := unitconv.Fahrenheit(t)
		c := unitconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, unitconv.FToC(f), c, unitconv.CToF(c))
	}

	{
		f := unitconv.Feet(t)
		m := unitconv.Meter(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, unitconv.FToM(f), m, unitconv.MToF(m))
	}

	{
		p := unitconv.Pound(t)
		k := unitconv.Kilogram(t)
		fmt.Printf("%s = %s, %s = %s\n",
			p, unitconv.PToK(p), k, unitconv.KToP(k))
	}

	{
		a := unitconv.Atm(t)
		b := unitconv.Bar(t)
		fmt.Printf("%s = %s, %s = %s\n",
			a, unitconv.AToB(a), b, unitconv.BToA(b))
	}
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			handle(arg)
		}
		return
	}

	fmt.Println("Input number. or Ctrl-C to quit.")

	for true {
		var arg string
		_, err := fmt.Scanf("%s", &arg)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		handle(arg)
	}
}

//!-
