/*
Exercise 2.2: Write a general-purpose unit-conversion program analogous to cf
that reads numbers from its command-line arguments or from the standard input
if there are no arguments, and converts each number into units like temperature
in Celsius and Fahrenheit, length in feet and meters, weight in pounds and kilograms, and the like.
*/

package unitconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

type Feet float64
type Meter float64

type Pound float64
type Kilogram float64

type Bar float64
type Atm float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%0.2f°C", c)
}
func (f Fahrenheit) String() string { return fmt.Sprintf("%0.2f°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%0.2f K", k) }
func (f Feet) String() string       { return fmt.Sprintf("%0.2f Feet", f) }
func (m Meter) String() string      { return fmt.Sprintf("%0.2f meter", m) }
func (p Pound) String() string      { return fmt.Sprintf("%0.2f pound", p) }
func (b Bar) String() string        { return fmt.Sprintf("%0.2f bar", b) }
func (a Atm) String() string        { return fmt.Sprintf("%0.2f atm", a) }
func (k Kilogram) String() string   { return fmt.Sprintf("%0.2f kg", k) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK coverts a Celsius temperature to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// KToC converts a Kelvin temperature to Celsius
func KToC(k Kelvin) Celsius { return Celsius(k - Kelvin(AbsoluteZeroC)) }

//FToM converts a Feet length to Meter
func FToM(f Feet) Meter { return Meter(f * 0.3048) }

//MToF converts a Meter length to Feet
func MToF(m Meter) Feet { return Feet(m / 0.3048) }

//PToK converts a Pound weight to Kilogram
func PToK(p Pound) Kilogram { return Kilogram(p * 0.453592) }

//KToP converts a Kilogram weight to Pound
func KToP(k Kilogram) Pound { return Pound(k / 0.453592) }

//AToB converts a Atm pressure to Bar
func AToB(a Atm) Bar { return Bar(a / 1.01325) }

//BToA converts a Bar pressure to Atm
func BToA(b Bar) Atm { return Atm(b * 1.01325) }

//!-
