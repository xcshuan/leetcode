package problem0003

import "testing"

func TestOk(t *testing.T) {
	var s = "qwererw"
	var expect = 4
	res := lengthOfLongestSubstring(s)
	if res != expect {
		t.Error("Test failed")
	}
}
