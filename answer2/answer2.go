package answer2

import "strconv"

func ContainsOnlyDigits(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
