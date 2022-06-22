package main

/**
 *
 * @param lists ListNode类一维数组
 * @return ListNode类
 */

func merge2(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	result := &ListNode{}
	cur := result
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val >= pHead2.Val {
			cur.Next = pHead2
			pHead2 = pHead2.Next
		} else {
			cur.Next = pHead1
			pHead1 = pHead1.Next
		}
		cur = cur.Next
	}

	if pHead1 == nil {
		cur.Next = pHead2
	} else if pHead2 == nil {
		cur.Next = pHead1
	}

	return result.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	// write code here
	if len(lists) < 1 {
		return nil
	}

	for len(lists) > 1 {
		for i := 0; i < len(lists)-1; i++ {
			lists[i] = merge2(lists[i], lists[len(lists)-1])
			lists[len(lists)-1] = nil
			lists = lists[:len(lists)-1]
		}
	}

	return lists[0]
}

func adjustUp(l *[]*ListNode, index int) {
	for index > 0 {
		if index&2 != 0 {
			if (*l)[(index-1)/2] == nil || ((*l)[index] != nil && (*l)[index].Val < (*l)[(index-1)/2].Val) {
				(*l)[index], (*l)[(index-1)/2] = (*l)[(index-1)/2], (*l)[index]
				adjustDown(l, index)
			} else {
				break
			}
			index = (index - 1) / 2
		} else {
			if (*l)[index/2] == nil || ((*l)[index] != nil && (*l)[index].Val < (*l)[index/2].Val) {
				(*l)[index], (*l)[index/2] = (*l)[index/2], (*l)[index]
				adjustDown(l, index)
			} else {
				break
			}
			index = index / 2
		}
	}

	return
}

// adjustDown
func pop(l *[]*ListNode) *ListNode {
	result := (*l)[0]
	if result == nil {
		return nil
	}
	(*l)[0] = (*l)[0].Next
	adjustDown(l, 0)
	return result
}

// adjustDown
func adjustDown(l *[]*ListNode, index int) {
	for index < len(*l) {
		if 2*index+1 < len(*l) && (*l)[2*index+1] != nil && ((*l)[index] == nil || (*l)[index].Val > (*l)[2*index+1].Val) {
			// 左不为nil
			(*l)[index], (*l)[2*index+1] = (*l)[2*index+1], (*l)[index]
			index = 2*index + 1
		} else if 2*index+2 < len(*l) && (*l)[2*index+2] != nil && ((*l)[index] == nil || (*l)[index].Val > (*l)[2*index+2].Val) {
			// 右不为nil
			(*l)[index], (*l)[2*index+2] = (*l)[2*index+2], (*l)[index]
			index = 2*index + 2
		} else {
			break
		}
	}
}

func mergeKLists2(lists []*ListNode) *ListNode {
	// write code here
	if len(lists) == 0 {
		return nil
	}

	for i := len(lists) - 1; i > 0; i-- {
		adjustUp(&lists, i)
	}
	result := &ListNode{}
	cur := result

	for {
		l := pop(&lists)
		if l == nil {
			break
		}
		cur.Next = l
		cur = cur.Next
	}

	return result.Next
}
