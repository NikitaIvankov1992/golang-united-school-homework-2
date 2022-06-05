package somename

import (
	"math"
)

func CalcSquare(sideLen float64, sidesNum int64) float64 {
	switch sidesNum {
	case 0:
		return 3.14 * (sideLen * sideLen)
	case 3:
		return 0.5 * sideLen * ((sideLen * math.Sqrt(3)) / 2)
	case 4:
		return sideLen * sideLen
	default:
		return 0
	}
}
