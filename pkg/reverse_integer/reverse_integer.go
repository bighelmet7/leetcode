package reverse_integer

import "math"

func totalDigits(x int) int {
	return int(math.Log10(float64(x))) + 1
}

func reverse(x int) int {
	// [-9, 9]
	if x >= -9 && x <= 9 {
		return x
	}
	var isNeg bool
	if x < 0 {
		isNeg = true
		x *= -1
	}
	digits := totalDigits(x)
	reverseNum := make([]int, digits)
	iter := 0
	for x > 0 {
		lastDigit := int(x % 10)
		reverseNum[iter] = lastDigit
		x /= 10
		iter++
	}
	var result int
	iter = 0
	for i := len(reverseNum) - 1; i >= 0; i-- {
		result += reverseNum[i] * int(math.Pow10(iter))
		iter++
	}
	if isNeg {
		result *= -1
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}
