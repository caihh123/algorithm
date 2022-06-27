package main

import (
	"testing"
)

func Test_DeleteDuplicates(t *testing.T) {
	arr := []int{
		1, 1, 2,
	}
	l1 := &ListNode{}
	cur1 := l1
	for _, num := range arr {
		cur1.Next = &ListNode{
			Val: num,
		}
		cur1 = cur1.Next
	}

	result := deleteDuplicates(l1.Next)
	t.Log(result)
}
