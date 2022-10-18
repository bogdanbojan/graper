package graper

import "math"

// monitorFerment monitors ferment from the refractometer readings.
func monitorFerment(ib float64, cb float64) {
	// Initial gravity.
	ig := 1.000898 + 0.003859118*ib + 0.00001370735*math.Pow(ib, 2) + 0.00000003742517*math.Pow(ib, 3)

	// Current gravity based on ib and cb.
	cg := 1.001843 - 0.002318474*ib - 0.000007775*math.Pow(ib, 2) - 0.000000034*math.Pow(ib, 3) + 0.00574*cb + 0.00003344*math.Pow(cb, 2) + 0.000000086*math.Pow(cb, 3)

	// Alcohol by volume.

}
