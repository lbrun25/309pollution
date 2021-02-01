package utils

import (
	"math"
)

func Factorial(n uint64)(result uint64) {
	if n > 0 {
		result = n * Factorial(n - 1)
		return result
	}
	return 1
}

func BinomialCoefficient(n uint64, i uint64) float64 {
	return float64(Factorial(n) / (Factorial(i) * Factorial(n - i)))
}

func BernsteinPolynomial(n uint64, i uint64, u float64) float64 {
	a := BinomialCoefficient(n, i)
	b := math.Pow(u, float64(i))
	c := math.Pow(1 - u, float64(n - i))
	return a * b * c
}