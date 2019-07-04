package problem0002

import "testing"

func ToListNode(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}

	res := &ListNode{
		Val: list[0],
	}
	temp := res

	for _, val := range list {
		temp.Next = &ListNode{Val: val}
		temp = temp.Next
	}

	return res
}

func TestAddTwoNumbers(t *testing.T) {
	list1 := []int{2, 4, 3}
	list2 := []int{5, 6, 4}
	l1 := ToListNode(list1)
	l2 := ToListNode(list2)

	addTwoNumbers(l1, l2)

	list3 := []int{5}
	list4 := []int{5}
	l3 := ToListNode(list3)
	l4 := ToListNode(list4)
	addTwoNumbers(l3, l4)
}
