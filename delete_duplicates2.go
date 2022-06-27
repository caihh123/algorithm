package main

func deleteDuplicates2(head *ListNode) *ListNode {
	// write code here

	result := &ListNode{
		Next: head,
	}
	pre := result
	cur := head

	for cur != nil {
		tempCur := cur.Next
		count := 0
		for tempCur != nil && tempCur.Val == cur.Val {
			count++
			tempCur = tempCur.Next
		}

		if count > 0 {
			pre.Next = tempCur
			cur = tempCur
		} else {
			pre = cur
			cur = cur.Next
		}
	}

	return result.Next
}
