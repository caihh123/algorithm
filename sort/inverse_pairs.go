package sort

func InversePairs(data []int) int {
	// write code here
	temp := make([]int, len(data))

	for i := len(data) - 2; i >= 0; i-- {
		for j := i + 1; j < len(data); j++ {
			if data[i] == data[j] {
				temp[i] = temp[j]
				break
			} else if data[i] > data[j] {
				temp[i] = temp[j] + 1
				break
			}
		}
	}

	var count int
	for _, num := range temp {
		count += num
	}

	return count % 1000000007
}
