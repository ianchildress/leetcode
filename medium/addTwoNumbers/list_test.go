package main

import (
	"testing"
)

func Test_newList(t *testing.T) {
	node := newList([]int{2, 4, 3})

	if node.Val != 3 {
		t.Errorf("wanted %v got %v", 3, node.Val)
	}
	if node.Next.Val != 4 {
		t.Errorf("wanted %v got %v", 4, node.Val)
	}
	if node.Next.Next.Val != 2 {
		t.Errorf("wanted %v got %v", 2, node.Val)
	}
}

func Test_length(t *testing.T) {
	list := []int{2, 4, 3}
	node := newList(list)
	if listLength(node) != len(list) {
		t.Errorf("wanted %v got %v", len(list), listLength(node))
	}
	list = []int{2, 4, 3, 6}
	node = newList(list)
	if listLength(node) != len(list) {
		t.Errorf("wanted %v got %v", len(list), listLength(node))
	}

	list = []int{2, 4, 3, 6, 9}
	node = newList(list)
	if listLength(node) != len(list) {
		t.Errorf("wanted %v got %v", len(list), listLength(node))
	}
}
