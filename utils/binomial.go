package utils

import (
	"errors"
	"math/big"
)

var (
	ErrFactorialNumber = errors.New("the number is way too big to calculate a factorial")
	ErrKGreaterThanN = errors.New("k > n")
)

// Factorial returns the factorial big int
func Factorial(x *big.Int) (*big.Int, error) {
	result := big.NewInt(1)
	i := big.NewInt(2)

	if !x.IsInt64() {
		return nil, ErrFactorialNumber
	}
	for i.Cmp(x) != 1 {
		result.Mul(result, i)
		i = i.Add(i, big.NewInt(1))
	}
	return result, nil
}

func getBinomialCoefficient(n *big.Int, k *big.Int) (*big.Int, error) {
	if k.Cmp(n) == 1 {
		return nil, ErrKGreaterThanN
	}

	numerator, err := Factorial(n); if err != nil {
		return nil, err
	}
	subNK := big.NewInt(1).Sub(n, k)
	fK, err := Factorial(k); if err != nil {
		return nil, err
	}
	fSubNK, err := Factorial(subNK); if err != nil {
		return nil, err
	}
	denominator := big.NewInt(1).Mul(fK, fSubNK)
	res := big.NewInt(1).Div(numerator, denominator)
	return res, nil
}

// BigPow returns big Float
func BigPow(a *big.Float, e int64) *big.Float {
	if e == 0 {
		return big.NewFloat(1.0)
	}
	result := big.NewFloat(0.0).Copy(a)
	for i := int64(0); i < e - 1; i++ {
		result = result.Mul(result, a)
	}
	return result
}

// BernsteinPolynomial returns the Bernstein basis polynomial
func BernsteinPolynomial(n int64, k int64, p float64) (float64, error) {
	binomialCoefficient, err := getBinomialCoefficient(big.NewInt(0.0).SetInt64(n), big.NewInt(0.0).SetInt64(k))
	if err != nil {
		return 0, err
	}
	res := big.NewFloat(0.0).Mul(
		big.NewFloat(0.0).SetInt(binomialCoefficient),
		BigPow(big.NewFloat(0.0).SetFloat64(p), k))
	res.Mul(res, BigPow(big.NewFloat(0.0).Sub(big.NewFloat(1.0), big.NewFloat(0.0).SetFloat64(p)), n - k))

	convertedRes, _ := res.Float64()
	return convertedRes, nil
}