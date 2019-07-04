package problem0001

import "fmt"

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range nums {
		if a, ok := numsMap[target-num]; ok {
			return []int{a, i}
		}
		numsMap[num] = i
	}
	return nil
}

func main() {
	var a = []int{2, 7, 11, 15}
	var b = 9
	res := twoSum(a, b)
	var expect = []int{0, 1}
	if false == equal(res, expect) {
		panic("Test failed")
	}
	fmt.Println("Test success!")
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
