package main

import (
	"fmt"
	"math/big"
)

//Given n points on a 2D plane, find the maximum number of points that lie on the same straight line.

/**
 * Definition for a point.
 */
type Point struct {
	X int
	Y int
}

func maxPoints(points []Point) int {
	if len(points) < 3 {
		return len(points)
	}

	var res int
	for i := 0; i < len(points); i++ {
		var record = make(map[string]int)
		var numSame, numVertical, local = 0, 1, 1
		for j := i + 1; j < len(points); j++ {
			if points[i].X == points[j].X {
				if points[i].Y == points[j].Y {
					numSame++
				} else {
					numVertical++
				}
			} else {
				r := big.NewRat(int64(points[i].Y-points[j].Y), int64(points[i].X-points[j].X))
				if _, ok := record[r.RatString()]; ok {
					record[r.RatString()]++
				} else {
					record[r.RatString()] = 2
				}
				if local < record[r.RatString()] {
					local = record[r.RatString()]
				}
			}
		}
		if local+numSame > numSame+numVertical {
			local += numSame
		} else {
			local = numSame + numVertical
		}
		if local > res {
			res = local
		}
	}

	return res
}

func main() {
	var points = []Point{
		//Point{84, 250}, Point{0, 0}, Point{1, 0},
		//Point{0, -70}, Point{0, -70}, Point{1, -1},
		//Point{21, 10}, Point{42, 90}, Point{-42, -230},
		//		Point{0, 0}, Point{-1, -1}, Point{2, 2},
		Point{0, 0}, Point{94911151, 94911150}, Point{94911152, 94911151},
	}
	fmt.Println(maxPoints(points))
}
