package missingnumber

// MissingNumber will find the missing number in nums such as nums goes from 0,2,..., n.
// nums doesn't have to be sorted
func MissingNumber(nums []int) int {
	var container int

	for i := 0; i <= len(nums); i++ {
		container ^= i
	}

	for _, num := range nums {
		container ^= num
	}

	return container
}
