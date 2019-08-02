package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(a, b *ListNode) *ListNode {
	root := a
	aptr := a.Next
	bptr := b
	sw := true
	for aptr != nil || bptr != nil {
		if sw { // increment node
			root.Next = aptr
			if aptr.Next != nil {
				aptr = aptr.Next
			}
		} else {
			root.Next = bptr
			if bptr.Next != nil {
				bptr = bptr.Next
			}
		}
		sw = !sw
	}
	return root
}
