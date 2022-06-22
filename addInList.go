package main

func reverse(pHead *ListNode) *ListNode {
	var (
		pre = pHead
		cur = pHead
	)

	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next = cur.Next.Next

		next.Next = pre
		pre = next
	}

	return pre
}

func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
	// write code here

	cur1 := reverse(head1)
	cur2 := reverse(head2)
	result := &ListNode{
		Next: cur1,
	}
	pre1 := result

	add := 0

	for cur1 != nil && cur2 != nil {
		cur1.Val += cur2.Val
		cur1.Val += add
		add = 0
		if cur1.Val >= 10 {
			cur1.Val %= 10
			add = 1
		}

		pre1 = cur1
		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	if cur2 != nil {
		pre1.Next = cur2
		cur1 = pre1
		cur1 = cur1.Next
	}

	for cur1 != nil {
		cur1.Val += add
		add = 0
		if cur1.Val >= 10 {
			cur1.Val %= 10
			add = 1
		}
		if cur1.Next == nil {
			break
		}
		cur1 = cur1.Next
	}

	if add != 0 {
		cur1.Next = &ListNode{
			Val: add,
		}
	}

	return reverse(result.Next)
}
