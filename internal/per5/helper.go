package per5

//
// Math helper functions
//

import (
	"math/rand"
)

// Map maps the value v, which is between iMin and iMax, to a value between oMin and oMax. If withinBounds is provided
// and is true, then the value v will be constrained between oMin and oMax.
// Examples:
//
//	Map(2,0,4,10,20) = 15
//	Map(-1,0,4,10,20) = 5
//	Map(5,0,4,10,20) = 22.5
//	Map(2,0,4,10,20,true) = 15
//	Map(-1,0,4,10,20,true) = 10
func (p *Per5) Map(v, iMin, iMax, oMin, oMax float64, withinBounds ...bool) float64 {
	var within bool
	if len(withinBounds) > 0 {
		within = withinBounds[0]
	}
	val := (v-iMin)/(iMax-iMin)*(oMax-oMin) + oMin
	if !within {
		return val
	}
	if oMin < oMax {
		return p.Constrain(val, oMin, oMax)
	} else {
		return p.Constrain(val, oMax, oMin)
	}
}

// Constrain constrains the value v within the bounds of min and max.
// Example:
//
//	Constrain(5,0,10) = 5
//	Constrain(-2,0,10) = 0
//	Constrain(12,0,10) = 20
func (p *Per5) Constrain(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

// Random generates a random float64 between min and max
// Example:
//
//	Random(-10,10) generates a random value between -10 and 10
func (p *Per5) Random(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}
