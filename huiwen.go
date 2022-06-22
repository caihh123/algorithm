package main

func isPail(head *ListNode) bool {
	// write code here
	if head == nil || head.Next == nil {
		return true
	}

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 反转
	var pre *ListNode
	cur := slow
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	mid := pre
	for mid != nil {
		if mid.Val != head.Val {
			return false
		}
		mid = mid.Next
		head = head.Next
	}

	return true
}
