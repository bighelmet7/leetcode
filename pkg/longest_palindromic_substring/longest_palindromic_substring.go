package longest_palindromic_substring

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	var result string
	memo := map[rune][]int{}
	for i := 0; i < len(s); i++ {
		memo[rune(s[i])] = append(memo[rune(s[i])], i)
	}
	// one single letter
	if len(memo) == 1 {
		return string(s)
	}
	for i := 0; i < len(s); i++ {
		var tmp string
		values := memo[rune(s[i])]
		// we can potentially have a palindrom
		if len(values) > 1 {
			for j := len(values) - 1; j >= 0; j-- {
				rightIndex := values[j]
				// different character
				if rightIndex > i {
					tmp = s[i : rightIndex+1]
					if isPalindrome(tmp) && len(tmp) > len(result) {
						result = tmp
					}
				}
			}
		}
	}
	if result == "" && len(s) > 0 {
		result = string(s[0])
	}
	return result
}
