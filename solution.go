package square

import (
	"math"
)

type NewType int

const SidesSquare NewType = 4
const SidesTriangle NewType = 3
const SidesCircle NewType = 0

func CalcSquare(sideLen float64, sidesNum NewType) float64 {
	switch sidesNum {
	case 0:
		return math.Pi * (sideLen * sideLen)
	case 3:
		return 0.5 * sideLen * ((sideLen * math.Sqrt(3)) / 2)
	case 4:
		return sideLen * sideLen
	default:
		return 0
	}
}
