package sort

func merge(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}

	mid := len(arr) >> 1
	left := arr[:mid]
	right := arr[mid:]

	return merge2(merge(left), merge(right))
}

func merge2(left, right []int) []int {

	result := make([]int, 0, len(left)+len(right))

	for len(left) != 0 && len(right) != 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	if len(left) != 0 {
		result = append(result, left...)
	}

	if len(right) != 0 {
		result = append(result, right...)
	}

	return result

}
