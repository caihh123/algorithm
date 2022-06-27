package sort

func Binary(nums []int, target int) int {
	// write code here
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (right + left) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}
