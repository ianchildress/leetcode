package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// newList returns a linked list from the given int slice
func newList(list []int) *ListNode {
	var lastNode *ListNode
	for i := range list {
		n := &ListNode{
			Val: list[i],
		}

		n.Next = lastNode
		lastNode = n
	}
	return lastNode
}

// listLength returns the number of does in the list
func listLength(node *ListNode)int{
	var i int
	for node!=nil{
		i++
		node=node.Next
	}
	return i
}
