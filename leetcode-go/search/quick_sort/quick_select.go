package search

func quickSelect(arr []int, k int) int {
	return quickSelectHelper(arr, 0, len(arr)-1, k)
}

func quickSelectHelper(arr []int, low, high, k int) int {
	if low == high {
		return arr[low]
	}
	p := partition(arr, low, high)
	if p == k {
		return arr[p]
	} else if p > k {
		return quickSelectHelper(arr, low, p-1, k)
	} else {
		return quickSelectHelper(arr, p+1, high, k)
	}
}
