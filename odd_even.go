package main

func oddEvenList(head *ListNode) *ListNode {
	// write code here
	odd := &ListNode{}
	curOdd := odd
	even := &ListNode{}
	curEven := even

	var count int
	for head != nil {
		next := head.Next
		head.Next = nil
		if count&1 == 0 {
			curOdd.Next = head
			curOdd = curOdd.Next
		} else {
			curEven.Next = head
			curEven = curEven.Next
		}
		count++
		head = next
	}

	curOdd.Next = even.Next
	return odd.Next
}
