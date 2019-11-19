/*
Exercise 2.1: Add types, constants, and functions to tempconv for processing temperatures in the Kelvin scale,
where zero Kelvin is −273.15°C and a difference of 1K has the same magnitude as 1°C.
*/

package main

import (
	"fmt"
	"github.com/jijiwhywhy/goplSolutions/chapter02/ex2.1/tempconv"
)

func main() {
	fmt.Println(tempconv.CToK(tempconv.FreezingC))
	fmt.Println(tempconv.CToK(tempconv.BoilingC))

}
