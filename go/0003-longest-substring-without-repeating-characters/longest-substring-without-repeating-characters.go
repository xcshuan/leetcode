package problem0003

func lengthOfLongestSubstring(s string) int {
	var ans int
	var n = len(s)
	charMap := make(map[byte]int)
	var i = 0
	for j := 0; j < n; j++ {
		if _, ok := charMap[s[j]]; ok {
			if charMap[s[j]] > i {
				i = charMap[s[j]]
			}
		}
		if j-i+1 > ans {
			ans = j - i + 1
		}
		charMap[s[j]] = j + 1
	}
	return ans
}
