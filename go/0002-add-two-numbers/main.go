// Definition for singly-linked list.
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var a, b, c, remain int

	var x, y = l1, l2
	ans := new(ListNode)
	temp := ans
	for x != nil || y != nil {
		a = 0
		b = 0
		if x != nil {
			a = x.Val
			x = x.Next
		}
		if y != nil {
			b = y.Val
			y = y.Next
		}

		c = (a + b + remain) % 10
		remain = (a + b + remain) / 10
		temp.Val = c
		if x != nil || y != nil {
			temp.Next = new(ListNode)
			temp = temp.Next
		}
	}
	if remain != 0 {
		temp.Next = new(ListNode)
		temp.Next.Val = remain
	}
	return ans
}

func main() {
	a := new(ListNode)
	a.Val = 5
	b := new(ListNode)
	b.Val = 3
	c := new(ListNode)
	c.Val = 9
	b.Next = c
	a.Next = b
	x := new(ListNode)
	x.Val = 3
	y := new(ListNode)
	y.Val = 7
	z := new(ListNode)
	z.Val = 6
	y.Next = z
	x.Next = y
	printListNode(a)
	printListNode(x)
	res := addTwoNumbers(a, x)
	printListNode(res)
}

func printListNode(l *ListNode) {
	var x = l
	fmt.Printf("[")
	for x != nil {
		fmt.Printf("%d", x.Val)
		if x.Next != nil {
			fmt.Printf(", ")
		}
		x = x.Next
	}
	fmt.Printf("]\n")
}
