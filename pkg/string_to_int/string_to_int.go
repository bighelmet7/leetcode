// #8 - String to int
package string_to_int

import "math"

func myAtoi(s string) int {
	positiveASCII := rune('+')
	negativeASCII := rune('-')
	zeroASCII := rune('0')
	nineASCII := rune('9')
	spaceASCII := rune(' ')

	var (
		isNegative bool
		isNumber   bool
		result     int
		prevChar   rune
	)
	for _, c := range s {
		// check if the current char does not match as a valid digit
		if prevChar == negativeASCII || prevChar == positiveASCII {
			if c < zeroASCII || c > nineASCII {
				break
			}
		}
		// if the previous char is a valid digit
		if prevChar != 0 && prevChar != spaceASCII && prevChar != negativeASCII && prevChar != positiveASCII {
			if prevChar < zeroASCII || prevChar > nineASCII {
				break
			}
		}
		if c == negativeASCII {
			isNegative = true
		}
		if c >= zeroASCII && c <= nineASCII {
			if !isNegative {
				result = (10 * result) + (int(c) - '0')
			} else {
				result = (10 * result) - (int(c) - '0')
			}
			// number too big
			if result > math.MaxInt32 {
				result = math.MaxInt32
				break
			}
			if result < math.MinInt32 {
				result = math.MinInt32
				break
			}
			isNumber = true
		} else if isNumber {
			break
		}
		prevChar = c
	}
	return result
}
