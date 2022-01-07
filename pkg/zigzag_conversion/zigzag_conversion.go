package zigzag_conversion

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	zigzag := make(map[int][]rune)
	var (
		iterator = 1
		diagonal bool
	)
	for i := 0; i < len(s); i++ {
		// diagonal case
		if iterator > numRows {
			iterator = numRows - 1
			diagonal = true
		}
		// we are back to columns
		if iterator == 1 && diagonal {
			diagonal = false
		}

		value, _ := zigzag[iterator]
		zigzag[iterator] = append(value, rune(s[i]))

		// reset values depending if you are on columns or rows
		if !diagonal {
			iterator++
		} else {
			iterator--
		}
	}
	var result string
	for i := 1; i <= numRows; i++ {
		result = result + string(zigzag[i][:])
	}
	return result
}
