package main

import (
	"testing"
)

func Test_Merge_k_List(t *testing.T) {

	arrs := [][]int{
		{-3, -3, -3}, {-9, -8, -7, -7, -6, -4, 2}, {-2, 3}, {-8, -4, -3, -2},
	}
	lists := make([]*ListNode, 0, len(arrs))
	for _, arr := range arrs {
		if len(arr) == 0 {
			lists = append(lists, nil)
			continue
		}
		l := &ListNode{Val: arr[0]}
		lists = append(lists, l)
		for i := 1; i < len(arr); i++ {
			l.Next = &ListNode{
				Val: arr[i],
			}
			l = l.Next
		}
	}

	lists2 := make([]*ListNode, 0, len(arrs))
	for _, arr := range arrs {
		if len(arr) == 0 {
			lists2 = append(lists2, nil)
			continue
		}
		l := &ListNode{Val: arr[0]}
		lists2 = append(lists2, l)
		for i := 1; i < len(arr); i++ {
			l.Next = &ListNode{
				Val: arr[i],
			}
			l = l.Next
		}
	}

	result1 := mergeKLists(lists)
	result2 := mergeKLists2(lists2)

	cur1 := result1
	cur2 := result2
	for cur1 != nil && cur2 != nil {
		if cur1.Val != cur2.Val {
			t.Error()
			return
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	if cur1 == nil || cur2 == nil {
		t.Error()
	}

}

func Test_addInList(t *testing.T) {
	// arr := [][]int{{5, 9, 2, 3, 7, 4, 9, 9, 0, 2, 6, 6, 1, 3, 8, 3, 2, 1, 9, 8, 4, 3, 1, 3, 3, 7, 5, 3, 9, 3, 1, 3, 1}, {4, 2, 8, 3, 5, 1, 0, 5, 7, 4, 5, 0, 2, 5, 0, 3, 9, 7, 3, 6, 8, 4, 4, 9, 7, 1}}
	arr := [][]int{
		{
			4, 6, 0, 2, 6, 6, 3, 6, 3, 0, 7, 8, 0, 4, 1, 7, 0, 5, 6, 5, 2, 4, 9, 9, 1, 5, 1, 5,
		},
		{
			6, 2, 7, 8, 6, 4, 7, 0, 9, 3, 0, 3, 6, 2, 5, 6, 0, 9, 6, 2, 7, 4, 2, 7, 1, 0, 9, 0, 5, 6, 5, 4, 9, 1, 8, 9, 3, 4, 0, 2, 1, 8, 8, 2, 2, 0,
		},
	}

	// 5, 9, 2, 3, 7, 4, 9, 9, 0, 2, 6, 6, 1, 3, 8, 3, 2, 1, 9, 8, 4, 3, 1, 3, 3, 7, 5, 3, 9, 3, 1, 3, 1
	//                      4, 2, 8, 3, 5, 1, 0, 5, 7, 4, 5, 0, 2, 5, 0, 3, 9, 7, 3, 6, 8, 4, 4, 9, 7, 1
	// 5, 9, 2, 3, 7, 4, 9, 9, 0, 2, 6, 6, 1, 3, 8, 3, 2, 1, 9, 8, 4, 3, 1, 3, 3, 7, 5, 3, 9, 8, 0, 0, 2

	l1 := &ListNode{}
	cur1 := l1
	for _, num := range arr[0] {
		cur1.Next = &ListNode{
			Val: num,
		}
		cur1 = cur1.Next
	}

	l2 := &ListNode{}
	cur2 := l2
	for _, num := range arr[1] {
		cur2.Next = &ListNode{
			Val: num,
		}
		cur2 = cur2.Next
	}
	result := addInList(l1, l2)
	t.Log(result)
}
