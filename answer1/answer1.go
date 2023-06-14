package answer1

func FindMissingElement(nums []int) int {
	n := len(nums)
	expectedSum := (n * (n + 1)) / 2
	actualSum := 0

	for _, num := range nums {
		actualSum += num
	}

	return expectedSum - actualSum
}
