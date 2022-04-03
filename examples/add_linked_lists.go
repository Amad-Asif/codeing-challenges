package examples

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	var newList *ListNode
	var newHead *ListNode
	var carry int
	var node1 = l1
	var node2 = l2
	for {
		if node1 == nil && node2 == nil {
			break
		}
		sum := carry
		carry = 0
		if node1 != nil {
			sum += node1.Val
		}
		if node2 != nil {
			sum += node2.Val
		}

		rem := sum % 10
		if x := sum / 10; x > 0 {
			carry = x
		}
		if newList == nil {
			newList = &ListNode{Val: rem, Next: nil}
			newHead = newList
		} else {
			newList.Next = &ListNode{Val: rem, Next: nil}
			newList = newList.Next
		}

		if node1 != nil && node1.Next == nil && node2 != nil && node2.Next == nil {
			break
		}

		if node1 != nil && node1.Next != nil {
			node1 = node1.Next
		} else {
			node1 = nil
		}

		if node2 != nil && node2.Next != nil {
			node2 = node2.Next
		} else {
			node2 = nil
		}
	}

	if carry != 0 {
		newList.Next = &ListNode{Val: carry, Next: nil}
	}

	return newHead
}
