package main

import (
	"testing"
)

func TestHuiWen(t *testing.T) {
	arr := []int{1, 2, 3, 2, 1}
	head := &ListNode{}
	cur := head

	for _, num := range arr {
		l := &ListNode{Val: num}
		cur.Next = l
		cur = cur.Next
	}
	result := isPail(head.Next)
	t.Log(result)

}

func TestReverse(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	head := &ListNode{}
	cur := head

	for _, num := range arr {
		l := &ListNode{Val: num}
		cur.Next = l
		cur = cur.Next
	}
	result := reverse(head.Next)
	t.Log(result)

}
