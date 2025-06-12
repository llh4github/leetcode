package search

import "slices"

func intersection(nums1 []int, nums2 []int) []int {

	slices.Sort(nums1)
	slices.Sort(nums2)
	pre := -1
	ans := make([]int, 0)

	m, n := len(nums1), len(nums2)
	for i, j := 0, 0; i < m && j < n; {

		if nums1[i] == nums2[j] {
			if nums1[i] != pre {
				ans = append(ans, nums1[i])
			}
			pre = nums1[i]
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return ans
}

func nextGreatestLetter(letters []byte, target byte) byte {

	n := len(letters)
	if letters[n-1] < target {
		return letters[0]
	}
	left, right := 0, n-1
	ans := n
	for left <= right {

		mid := left + (right-left)/2
		switch {
		case letters[mid] == target:
			left++
		case letters[mid] < target:
			left = mid + 1
		case letters[mid] > target:
			ans = min(ans, mid)
			right--
		}
	}
	if ans != n {
		return letters[ans]
	}
	return letters[0]
}
func arrangeCoins(n int) int {

	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		if mid*(mid+1) <= 2*n {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left

}
