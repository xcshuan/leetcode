package problem0004

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	var m = len(nums1)
	var n = len(nums2)
	if m > n {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}

	var iMin, iMax, halfLen = 0, m, (m + n + 1) / 2
	for iMin <= iMax {
		var i = (iMin + iMax) / 2
		var j = halfLen - i
		if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			var maxLeft = 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				if nums1[i-1] > nums2[j-1] {
					maxLeft = nums1[i-1]
				} else {
					maxLeft = nums2[j-1]
				}
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			var minRight = 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				if nums2[j] < nums1[i] {
					minRight = nums2[j]
				} else {
					minRight = nums1[i]
				}
			}

			return float64(maxLeft+minRight) / 2.0
		}
	}
	return 0
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var len1, i = len(nums1), 0
	var len2, j = len(nums2), 0
	var flag, b = 0, false
	if 0 == (len1+len2)%2 {
		flag = 1
	}
	var ans float64
	for k := 0; k < len1+len2; k++ {
		if b {
			break
		}
		if j == len2 || (i < len1 && j < len2 && nums1[i] < nums2[j]) {
			if flag == 1 && (k == (len1+len2)/2 || k == (len1+len2)/2-1) {
				ans += float64(nums1[i]) / 2.0
				if k == (len1+len2)/2 {
					b = true
				}
			} else if flag == 0 && k == (len1+len2)/2 {
				ans = float64(nums1[i])
				b = true
			}
			i++
			continue
		}
		if i == len1 || (i < len1 && j < len2 && nums2[j] <= nums1[i]) {
			if flag == 1 && (k == (len1+len2)/2 || k == (len1+len2)/2-1) {
				ans += float64(nums2[j]) / 2.0
				if k == (len1+len2)/2 {
					b = true
				}
			} else if flag == 0 && k == (len1+len2)/2 {
				ans = float64(nums2[j])
				b = true
			}
			j++
		}
	}

	return ans
}
