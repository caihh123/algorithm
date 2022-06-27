package sort

func Find(target int, array [][]int) bool {
	// write code here
	rows := len(array)
	cols := len(array[0])
	found := false
	if rows > 0 && cols > 0 {
		row := 0
		col := cols - 1
		for row < rows && col >= 0 {
			if array[row][col] == target {
				found = true
				break
			} else if array[row][col] > target {
				col--
			} else {
				row++
			}
		}
	}
	return found
}
