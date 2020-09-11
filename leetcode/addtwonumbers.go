// https://leetcode.com/problems/add-two-numbers/
package leetcode

import (
	"fmt"
	"strconv"
)

// ListNode is meant to be a linked list representing a number. Each
// linked list node will hold a single digit and the digits are stored
// in reverse order (i.e. 2 -> 5 -> 1 is the number 152).
type ListNode struct {
	Val  int
	Next *ListNode
}

// String returns what the ListNode looks like. Useful for displaying
// this value nicely in test cases.
func (l ListNode) String() string {
	s := strconv.Itoa(l.Val)
	listIterator := l.Next
	for listIterator != nil {
		s = fmt.Sprintf("%s->%s", s, strconv.Itoa(listIterator.Val))
		listIterator = listIterator.Next
	}
	return s
}

// NumToListNode converts a number to a ListNode
func NumToListNode(n int) *ListNode {
	list := &ListNode{}
	listIterator := list
	quot := n
	for {
		rem := quot % 10
		quot = quot / 10
		listIterator.Val = rem
		if quot == 0 {
			break
		}
		listIterator.Next = &ListNode{}
		listIterator = listIterator.Next
	}
	return list
}

// listNodeToNum converts a ListNode to a number
func listNodeToNum(l *ListNode) int {
	multiplier := 1
	n := 0
	for {
		n += multiplier * l.Val
		if l.Next == nil {
			break
		}
		multiplier *= 10
		l = l.Next
	}
	return n
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{}
	listIterator := list
	carry := 0
	for {
		l1Val := 0
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}
		l2Val := 0
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}
		sum := l1Val + l2Val + carry
		carry = sum / 10
		digit := sum % 10
		listIterator.Val = digit
		if carry == 0 && l1 == nil && l2 == nil {
			break
		}
		listIterator.Next = &ListNode{}
		listIterator = listIterator.Next
	}
	return list
}

// A cute approach where I just convert the values to numbers and do
// the addition
func AddTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := listNodeToNum(l1)
	n2 := listNodeToNum(l2)
	return NumToListNode(n1 + n2)
}
