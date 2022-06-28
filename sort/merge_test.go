package sort

import (
	"testing"
)

func Test_Merge(t *testing.T) {
	arr := []int{
		1, 2, 4, 3, 5, 7, 6, 8, 9,
	}
	result := merge(arr)
	t.Log(result)
}
