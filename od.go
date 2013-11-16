// od is a command-line tool to calculate the necessary optical
// density of laser eye protection for continuous wave diode lasers
// with a wavelength between 315-1400nm. For reference, consider that
// a blue diode laser is typically around 445nm, a green diode laser
// is typically around 532nm, and a red diode laser is typically around
// 808nm.
package main

import (
	"flag"
	"math"
	"fmt"
)

var (
	L1 = math.Pow10(2)
	L2 = math.Pow10(3)
	L3 = math.Pow10(4)
	L4 = math.Pow10(5)
	L5 = math.Pow10(6)
	L6 = math.Pow10(7)
	L7 = math.Pow10(8)
	L8 = math.Pow10(9)
	L9 = math.Pow10(10)
	L10 = math.Pow10(11)
)

func powerDensity(beamArea, power float64) float64 {
	return power / beamArea
}

func beamArea(diameter float64) float64 {
	r := diameter / 2
	return math.Pi * (r * r)
}

func opticalDensity(powerDensity float64) int {
	switch {
	case powerDensity > L10:
		return 11
	case powerDensity > L9:
		return 10
	case powerDensity > L8:
		return 9
	case powerDensity > L7:
		return 8
	case powerDensity > L6:
		return 7
	case powerDensity > L5:
		return 6
	case powerDensity > L4:
		return 5
	case powerDensity > L3:
		return 4
	case powerDensity > L2:
		return 3
	case powerDensity > L1:
		return 2
	case powerDensity <= L1:
		return 1
	default:
		return -1
	}
}

func main() {
	diameter := flag.Float64("d", 9.0, "beam diameter in millimetres")
	power := flag.Float64("p", 1.0, "laser power in watts")
	flag.Parse()

	if *diameter <= 0.0 {
		fmt.Println("Beam diameter must be a positive number greater than zero.")
		return
	}

	*diameter /= 1000.0 // convert mm to m
	ba := beamArea(*diameter)
	pd := powerDensity(ba, *power)
	od := opticalDensity(pd)
	fmt.Printf("Beam area: %f m^2\nPower density: %f W/m^2\n", ba, pd)
	if od == -1 {
		fmt.Println("Couldn't determine the appropriate optical density rating.")
		fmt.Println("Please consult EN207 for the appopriate PPE.")
	} else if od == 11 {
		fmt.Printf("Your optical density requirements are above the ")
		fmt.Printf("standard requirements. Please consult EN207.\n")
	} else {
		fmt.Printf("Based on your laser's power and beam diameter, ")
		fmt.Printf("you should use eyewear with an OD of %d.\n", od)
	}
}
