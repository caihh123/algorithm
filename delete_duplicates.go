package main

func deleteDuplicates(head *ListNode) *ListNode {
	// write code here
	cur := head
	pre := head

	for cur != nil {
		if pre != nil && pre.Val == cur.Val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}

	return head
}
