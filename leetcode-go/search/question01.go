package search

import (
	"math/rand"
	"slices"
	"sort"
)

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

// 这题难在数量关系上
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sum := 0
	m := make(map[int]struct{})
	for _, v := range aliceSizes {
		sum += v
		m[v] = struct{}{}
	}

	for _, v := range bobSizes {
		sum -= v
	}

	delta := sum / 2
	for _, v := range bobSizes {
		tmp := v + delta
		if _, ok := m[tmp]; ok {
			return []int{tmp, v}
		}
	}
	return []int{}
}

type pair struct {
	indx  int
	count int
}

// 普通解法
func kWeakestRows(mat [][]int, k int) []int {

	tmp := make([]pair, 0)
	for i, row := range mat {
		cnt := 0
		for _, v := range row {
			if v == 0 {
				break
			} else {
				cnt++
			}
		}
		tmp = append(tmp, pair{i, cnt})
	}
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].count == tmp[j].count {
			return tmp[i].indx < tmp[j].indx
		}
		return tmp[i].count < tmp[j].count
	})
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = tmp[i].indx
	}
	return ans
}

func findKthLargest(nums []int, k int) int {
	return findKthLargestHelper(nums, 0, len(nums)-1, k-1)
}
// leetcode 构选的测试集会运行超时
func findKthLargestHelper(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}
	pIndex := partition(nums, left, right)
	if k == pIndex {
		return nums[k]
	} else if k < pIndex {
		return findKthLargestHelper(nums, left, pIndex-1, k)
	} else {
		return findKthLargestHelper(nums, pIndex+1, right, k)
	}

}
func partition(arr []int, low, high int) int {
	rIdx := rand.Intn(high-low+1) + low
	arr[high], arr[rIdx] = arr[rIdx], arr[high]
	pivot := arr[high]

	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] >= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
