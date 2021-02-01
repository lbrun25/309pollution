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

func buildMatrix() [][]int64 {
	matrix := make([][]int64, N)
	for i := 0; int64(i) < N; i++ {
		matrix[i] = make([]int64, N)
	}

	for _, point := range Points {
		matrix[point.X][point.Y] = point.P
	}
	return matrix
}

func bezierSurface(matrix [][]int64) (float64, error) {
	n := N - 1
	x := X / float64(n)
	y := Y / float64(n)
	res := 0.0

	for i := int64(0); i < n + 1; i++ {
		for j := int64(0); j < n + 1; j++ {
			bernX, err := utils.BernsteinPolynomial(n, i, x); if err != nil {
				return 0.0, err
			}
			bernY, err := utils.BernsteinPolynomial(n, j, y); if err != nil {
				return 0.0, err
			}
			res += bernX * bernY * float64(matrix[i][j])
		}
	}
	return res, nil
}

// Main - 309pollution main
func Main() error {
	matrix := buildMatrix()
	result, err := bezierSurface(matrix); if err != nil {
		return err
	}
	fmt.Printf("%.2f\n", result)
	return nil
}