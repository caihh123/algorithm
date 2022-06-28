package sort

func findPeakElement(nums []int) int {
	// write code here

	if len(nums) <= 1 {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			if i >= len(nums)-1 || nums[i] > nums[i+1] {
				return i
			}
		} else {
			i++
		}
	}

	return 0
}
