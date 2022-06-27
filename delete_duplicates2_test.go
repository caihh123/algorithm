package main

import (
	"testing"
)

func Test_DeleteDuplicates2(t *testing.T) {
	arr := []int{
		1, 1, 1, 2, 3,
	}
	l1 := &ListNode{}
	cur1 := l1
	for _, num := range arr {
		cur1.Next = &ListNode{
			Val: num,
		}
		cur1 = cur1.Next
	}

	result := deleteDuplicates2(l1.Next)
	t.Log(result)
}
