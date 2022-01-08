package palindrome_number

func isPalindrome(x int) bool {
	// one digit is always a palindrome
	if x >= 0 && x <= 9 {
		return true
	}
	// negatives or numbers that ends with 0
	// never will be palindromes
	if x < 0 || x%10 == 0 {
		return false
	}
	var result int
	for tmp := x; tmp != 0; tmp /= 10 {
		result = (result * 10) + (tmp % 10)
	}
	return result == x
}
