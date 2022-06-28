package sort

import (
	"testing"
)

func Test_findPeakElement(t *testing.T) {
	arr := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	index := findPeakElement(arr)
	t.Log(index)
}
