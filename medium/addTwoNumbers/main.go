package main

import "fmt"

/*
[9,9,9,9,9,9,9]
[9,9,9,9]

 */

func main() {
	l1 := newList([]int{9,9,9,9,9,9,9})
	l2 := newList([]int{9,9,9,9})
	n := addTwoNumbers(l1, l2)
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var head, tail *ListNode

	for l1 != nil || l2 != nil {
		var i1, i2 int
		if l1 != nil {
			i1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			i2 = l2.Val
			l2 = l2.Next
		}

		// add this nodes value
		sum := i1 + i2 + carry

		// check for carry
		carry = int(sum / 10) // 11 / 10 = 1

		// find value in 1's column
		val := sum % 10 // the remainder when dividing by 10 is what is in the 1's column

		// check if we initialize the head
		if head == nil {
			head = &ListNode{
				Val: val,
			}
			tail = head
			continue
		}

		// add a new tail
		tail.Next = &ListNode{
			Val: val,
		}
		tail = tail.Next
	}

	// check for extra carry
	if carry > 0 {
		tail.Next = &ListNode{
			Val: carry,
		}
	}

	return head
}
