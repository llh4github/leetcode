package search

import "math/rand"

func partition(arr []int, low, high int) int {
	rIdx := rand.Intn(high-low+1) + low
	// rIdx := high
	arr[high], arr[rIdx] = arr[rIdx], arr[high]
	pivot := arr[high]

	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
func quickSort(arr []int) {

	if len(arr) <= 1 {
		return
	}
	helper(arr, 0, len(arr)-1)
}

func helper(arr []int, low, high int) {
	if low < high {
		pIdx := partition(arr, low, high)
		helper(arr, low, pIdx-1)
		helper(arr, pIdx+1, high)
	}
}
