package pollution

import (
	"309pollution/utils"
	"fmt"
)

type Point struct {
	X int64
	Y int64
	P int64
}

// Values retrieved from the parser
var (
	N int64
	X float64
	Y float64
	Points []Point
)

func bezierSurface() float64 {
	n := N - 1
	x := X / float64(n)
	y := Y / float64(n)
	res := 0.0

	for i := int64(0); i < n + 1; i++ {
		for j := int64(0); j < n + 1; i++ {
			bernX := utils.GetBinomial(n, i, x)
			bernY := utils.GetBinomial(n, j, y)

			res += bernX * bernY * float64(Points[i].P)
		}
	}
	return res
}

func displayResult() {
	fmt.Printf("%.2f", bezierSurface())
}

// Main - 309pollution main
func Main() error {
	displayResult()
	return nil
}