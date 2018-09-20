package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

//时间复杂度：O(max(m,n))
//空间复杂度：O(max(m,n))
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//2个加数和进位
	var n1, n2, carry int
	l3 := &ListNode{}
	head := l3
	listInit := true
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
		}
		sum := n1 + n2 + carry
		n3 := sum % 10
		carry = sum / 10
		if listInit {
			head.Val = n3
		} else {
			l3.Next = &ListNode{
				Val:n3,
			}
			l3 = l3.Next
		}
		listInit = false
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return head

}

func main() {
	l11 := &ListNode{
		Val: 2,
		Next:&ListNode{
			Val:4,
			Next:&ListNode{
				Val:3,
			},
		},
	}
	l21 :=&ListNode{
		Val: 5,
		Next:&ListNode{
			Val:6,
			Next:&ListNode{
				Val:4,
			},
		},
	}
	l3 := addTwoNumbers(l11, l21)
	fmt.Println(l3.Val)
	fmt.Println(l3.Next.Val)
	fmt.Println(l3.Next.Next.Val)
}
