package search

import "testing"

func Test_findKthLargest(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		k    int
		want int
	}{
		{"test1", []int{3, 2, 1, 5, 6, 4}, 2, 5},
		{"test2", []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findKthLargest(tt.nums, tt.k)
			if tt.want != got {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
