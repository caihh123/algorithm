package sort

import (
	"testing"
)

func TestHeapArr(t *testing.T) {
	arr := []int{
		1, 5, 4, 9, 7, 6,
	}
	result := heapSort(arr)
	t.Log(result)
}
