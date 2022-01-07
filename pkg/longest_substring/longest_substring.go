package longest_substring

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var (
		subString    string
		maxSubString int
	)
	for i := 0; i < len(s); i++ {
		c := fmt.Sprintf("%c", s[i])
		if !strings.Contains(subString, c) {
			subString = subString + c
			maxSubString = max(maxSubString, len(subString))
		} else {
			j := strings.Index(subString, c)
			subString = subString[j+1:]
			subString = subString + c
		}
	}
	return maxSubString
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
