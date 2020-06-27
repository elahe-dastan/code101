package main

import "fmt"

func main() {
	l1 := &ListNode {
		2,
		&ListNode{
			4,
			&ListNode{
				3,
				nil,
			},
		},
	}

	l2 := &ListNode {
		5,
		&ListNode{
			6,
			&ListNode{
				4,
				nil,
			},
		},
	}

	r := addTwoNumbers(l1, l2)
	for {
		if r == nil {
			break
		}

		fmt.Print(r.Val)
		r = r.Next
	}
}

type ListNode struct {
	Val int
	Next *ListNode
}

var carry = 0
var first = true

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tail := &ListNode{}

	for {
		if l1 == nil && l2 == nil {
			break
		}else {
			if first {
				result = addDigit(l1, l2)
				result.Next = tail
				first = false
			}else {
				tail = addDigit(l1, l2)
				tail = tail.Next
			}

			if l1 != nil {
				l1 = l1.Next
			}
			if l2 != nil {
				l2 = l2.Next
			}
		}
	}

	return result
}

func addDigit(l1 *ListNode, l2 *ListNode) *ListNode {
	result := ListNode{}

	v1 := 0
	v2 := 0

	if l1 != nil {
		v1 = l1.Val
	}
	if l2 != nil {
		v2 = l2.Val
	}

	n := v1 + v2 + carry

	if n < 10 {
		result.Val = n
		carry = 0
	}else {
		result.Val = n % 10
		carry = 1
	}

	return &result
}