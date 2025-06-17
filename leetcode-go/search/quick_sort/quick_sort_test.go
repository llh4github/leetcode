package search

import "testing"

func Test_partitionQuickSort(t *testing.T) {
	arr := []int{3, 2, 1, 5, 6, 4, 7}
	p := partition(arr, 0, len(arr)-1)
	t.Log(arr)
	t.Log(p)
}

func Test_quickSort(t *testing.T) {
	arr := []int{3, 2, 1, 5, 6, 4, 7}
	quickSort(arr)
	t.Log(arr)
}
