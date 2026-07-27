package _3451_3460

import (
	"math"
)

// leetcode problem No. 3453
func separateSquares(squares [][]int) float64 {
	maxY, totalArea := 0.0, 0.0
	for _, sq := range squares {
		y, l := sq[1], sq[2]
		totalArea += float64(l * l)
		if float64(y+l) > maxY {
			maxY = float64(y + l)
		}
	}

	// get total area blow the line and total area above the line
	getArea := func(targetY float64) (float64, float64) {
		areaAbove := 0.0
		areaBelow := 0.0
		for _, sq := range squares {
			y, sideLength := sq[1], sq[2]
			if float64(y+sideLength) < targetY {
				areaBelow += float64(sideLength * sideLength)
			} else if float64(y) > targetY {
				areaAbove += float64(sideLength * sideLength)
			} else {
				areaAbove += (float64(y+sideLength) - targetY) * float64(sideLength)
				areaBelow += (targetY - float64(y)) * float64(sideLength)
			}
		}

		return areaAbove, areaBelow
	}

	lo, hi := 0.0, maxY
	eps := 1e-5
	for math.Abs(hi-lo) > eps {
		mid := (hi + lo) / 2.0
		areaAbove, areaBelow := getArea(mid)
		if areaAbove <= areaBelow { // we need to find the minimum, so there should be less or equal
			hi = mid
		} else {
			lo = mid
		}
	}
	return hi
}
