package main

import (
	"testing"
)

func Test_OddEvenList(t *testing.T) {
	arr := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	l1 := &ListNode{}
	cur1 := l1
	for _, num := range arr {
		cur1.Next = &ListNode{
			Val: num,
		}
		cur1 = cur1.Next
	}

	result := oddEvenList(l1.Next)
	t.Log(result)
}
