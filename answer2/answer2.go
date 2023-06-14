package answer2

import "strconv"

func ContainsOnlyDigits(str string) bool {
	for _, c := range str {
		_, err := strconv.Atoi(string(c))
		if err != nil {
			return false
		}
	}

	return true
}
