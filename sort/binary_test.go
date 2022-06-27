package sort

import (
	"testing"
)

func TestBinary(t *testing.T) {
	nums := []int{
		-1, 0, 3, 4, 6, 10, 13, 14,
	}
	result := Binary(nums, 13)
	t.Log(result)
}
