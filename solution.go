package square

import (
	"math"
)

type sides int

const (
	SidesCircle   sides = 0
	SidesTriangle sides = 3
	SidesSquare   sides = 4
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum sides) float64 {

	if sidesNum == 3 {
		return math.Sqrt(3) / 4 * sideLen * sideLen
	}

	if sidesNum == 4 {
		return sideLen * sideLen
	}

	if sidesNum == 0 {
		return math.Pi * sideLen * sideLen
	}
	return 0
}
