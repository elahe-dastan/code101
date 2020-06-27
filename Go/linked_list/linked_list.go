package main

import "fmt"

func main() {
	l1 := &ListNode {
		5,
		nil,
	}

	l2 := &ListNode {
		5,
		nil,
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	head := result

	for {
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}else {
			result.Next = addDigit(l1, l2)
			result = result.Next

			if l1 != nil {
				l1 = l1.Next
			}
			if l2 != nil {
				l2 = l2.Next
			}
		}
	}

	return head.Next
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