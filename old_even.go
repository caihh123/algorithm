package main

func oddEvenList(head *ListNode) *ListNode {
	// write code here

	result := &ListNode{}
	result.Next = head
	cur := result.Next

	for cur != nil {
		cur = cur.Next
	}

	return result.Next
}
