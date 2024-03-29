package RomanToInteger

func RomanToInt(s string) int {
	var helper, result, total int

	romanMap := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := len(s) - 1; i >= 0; i-- {
		total = romanMap[s[i]]
		if total < helper {
			result -= total
		} else {
			result += total
		}
		helper = total
	}

	return result
}
