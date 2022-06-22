package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// write code here
	result := &ListNode{Next: head}
	pre := result
	cur1 := pre.Next
	cur2 := pre.Next
	for i := 0; i < n; i++ {
		cur2 = cur2.Next
	}
	for cur2 != nil {
		pre = cur1
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	pre.Next = cur1.Next

	return result.Next
}

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil || pHead2 == nil {
		return nil
	}

	cur1, cur2 := pHead1, pHead2
	for cur1 != cur2 {
		if cur1 == nil {
			cur1 = pHead2
		} else {
			cur1 = cur1.Next
		}

		if cur2 == nil {
			cur2 = pHead1
		} else {
			cur2 = cur2.Next
		}
	}
	return cur1
}
