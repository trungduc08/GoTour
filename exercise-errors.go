package main

import (
	"fmt"
	"math"
)

// const EPSILON = 0.0000000001

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {

	return fmt.Sprintf("cannot Sqrt negative number:%v", float64(e))
}
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		fmt.Println("hello")
		return 0, ErrNegativeSqrt(x)
	}
	return CalSqrt(x), nil
}
func CalSqrt(x float64) float64 {
	z := float64(1)
	for math.Abs(z*z-x) >= EPSILON {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

// func main() {
// 	fmt.Println(Sqrt(-8))
// }
