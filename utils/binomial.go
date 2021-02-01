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

// Or you can use this second simpler solution to compute Bernstein polynomial without using math/big

//func Factorial(n int64) (int64, error) {
//	if n >= 100 {
//		return 0, ErrFactorialNumber
//	}
//	if n > 0 {
//		f, err := Factorial(n - 1); if err != nil {
//			return 0, err
//		}
//		result := n * f
//		return result, nil
//	}
//	return 1, nil
//}
//
//func BinomialCoefficient(n int64, i int64) (float64, error) {
//	fN, err := Factorial(n); if err != nil {
//		return 0, err
//	}
//	fI, err := Factorial(i); if err != nil {
//		return 0, err
//	}
//	fNi, err := Factorial(n - i); if err != nil {
//		return 0, err
//	}
//	return float64(fN / (fI * fNi)), nil
//}
//
//func BernsteinPolynomial(n int64, i int64, u float64) (float64, error) {
//	a, err := BinomialCoefficient(n, i); if err != nil {
//		return 0, err
//	}
//	b := math.Pow(u, float64(i))
//	c := math.Pow(1 - u, float64(n - i))
//	res := a * b * c
//	return res, nil
//}