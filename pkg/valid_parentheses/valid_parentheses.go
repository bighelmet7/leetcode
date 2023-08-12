package valid_parentheses

func isValid(s string) bool {
    if len(s) == 1 {
        return false
    }

	opening := []rune{}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			opening = append(opening, c)
		} else {
            // nothing has been opened
            if len(opening) == 0 {
                return false
            }
            top := opening[len(opening) - 1]
            roundBrackets := top == '(' && c == ')'
            squareBrackets := top == '[' && c == ']'
            curlyBrackets := top == '{' && c == '}'
            if roundBrackets || squareBrackets || curlyBrackets {
                opening = opening[:len(opening) - 1]
            } else {
                return false
            }
		}
	}

	return len(opening) == 0
}
