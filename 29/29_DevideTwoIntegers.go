package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	sf := 1
	if dividend^divisor < 0 {
		sf = -1
	}
	abs := func(x int) int {
		if x < 0 {
			return -1 * x
		}
		return x
	}
	dividend = abs(dividend)
	divisor = abs(divisor)

	target := 0
	n := 0
	for dividend >= divisor {
		divisor <<= 1
		n++
	}
	divisor = divisor >> 1
	for i := 0; i < n; i++ {
		if dividend >= divisor {
			dividend -= divisor
			target = (target << 1) + 1
		} else {
			target <<= 1
		}
		divisor >>= 1
	}
	result := target * sf
	if result > math.MaxInt32 {
		return math.MaxInt32
	}
	if result < math.MinInt32 {
		return math.MinInt32
	}
	return result

}

func main() {

	fmt.Println(divide(-2147483648, -1) == 2147483647)
	fmt.Println(divide(-1, -1) == 1)
	fmt.Println(divide(3, 1) == 3)
	fmt.Println(divide(0, 100) == 0)
	fmt.Println(divide(10, 3) == 3)
	fmt.Println(divide(7, -3) == -2)
	fmt.Println(divide(5, 2) == 2)
	fmt.Println(divide(3, 1))
	fmt.Println(divide(10, 3))
	fmt.Println(divide(7, -3))
	fmt.Println(divide(5, 2))

}
