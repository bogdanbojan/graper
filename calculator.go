package main

import (
	"fmt"
	"math"
)

type ferment struct {
	initial_gravity    float64
	current_gravity    float64
	current_gravity_bh float64
	true_brix          float64
	residual_sugar     float64
	current_alchol     float64
}

func main() {
	ferment := monitorFerment(10, 23)
	fmt.Print(ferment)
}

// monitorFerment monitors ferment from the refractometer readings.
func monitorFerment(ib float64, cb float64) ferment {
	ig := 1.000898 + 0.003859118*ib + 0.00001370735*math.Pow(ib, 2) + 0.00000003742517*math.Pow(ib, 3)
	cg := 1.001843 - 0.002318474*ib - 0.000007775*math.Pow(ib, 2) - 0.000000034*math.Pow(ib, 3) + 0.00574*cb + 0.00003344*math.Pow(cb, 2) + 0.000000086*math.Pow(cb, 3)
	abv := 0.93 * ((1017.5596 - (277.4 * cg) + (1.33302+0.001427193*cb+0.000005791157*math.Pow(cb, 2))*((937.8135*(1.33302+0.001427193*cb+0.000005791157*math.Pow(cb, 2)))-1805.1228)) * (cg / 0.794))
	bh := 143.254*cg*cg*cg - 648.670*cg*cg + 1125.805*cg - 620.389
	si := (2*math.Sqrt(626159497)*math.Sqrt(35209254016727200*abv+448667639342033000) - 33520822512398) / 841662180975
	remsg := cg - (1 - (si / 1000)) + 1
	tb := 143.254*remsg*remsg*remsg - 648.670*remsg*remsg + 1125.805*remsg - 620.389
	rs := tb * cg * 10

	a := math.Round(ig*10000) / 10000
	b := math.Round(cg*10000) / 10000
	c := math.Round(rs*1) / 1
	d := math.Round(abv*10) / 10
	e := math.Round(bh*10) / 10
	f := math.Round(tb*10) / 10

	return ferment{
		initial_gravity:    a,
		current_gravity:    b,
		current_gravity_bh: c,
		true_brix:          d,
		residual_sugar:     e,
		current_alchol:     f,
	}

}

// SO2 Aspiration/Oxidation calculation.
func SO2(mnaoh float64, tnaoh float64, vol float64) float64 {
	so := (tnaoh * mnaoh * 1.6 * 1000 * 20) / vol

	return math.Round(so*100) / 100
}

// Titrable Acidity.
func TACalc(mnaoh float64, tnaoh float64, vol float64) float64 {
	ta := (tnaoh * mnaoh * 75) / vol

	return math.Round(ta*100) / 100
}

// Simple Deacification.
// TODO: Ask if we use 99% of the time calcium carbonate.
func DeAcid(cta float64, tta float64, vol float64) float64 {
	factor := 100.087 / 150.087
	mass := (((cta - tta) * vol) * factor) / 1000

	return math.Round(mass*1000) / 1000
}
