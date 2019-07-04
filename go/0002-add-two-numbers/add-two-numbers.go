// Definition for singly-linked list.
package problem0002

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
