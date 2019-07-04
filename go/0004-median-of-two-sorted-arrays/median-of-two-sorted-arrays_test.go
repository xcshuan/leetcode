package problem0004

import "testing"

func TestOk(t *testing.T) {
	a := []int{1, 2}
	b := []int{3, 4}
	var expect = 2.5
	res := findMedianSortedArrays(a, b)
	if res != expect {
		t.Error("test failed", res, expect)
	}

	a = []int{1, 3}
	b = []int{2}
	expect = 2
	res = findMedianSortedArrays(a, b)
	if res != expect {
		t.Error("test failed", res, expect)
	}
}
