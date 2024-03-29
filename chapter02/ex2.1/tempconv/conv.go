// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

/*
Exercise 2.1: Add types, constants, and functions to tempconv for processing temperatures in the Kelvin scale,
where zero Kelvin is −273.15°C and a difference of 1K has the same magnitude as 1°C.
*/

package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK coverts a Celsius temperature to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// KToC converts a Kelvin temperature to Celsius
func KToC(k Kelvin) Celsius { return Celsius(k - Kelvin(AbsoluteZeroC)) }

//!-
