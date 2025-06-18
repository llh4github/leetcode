package search

import "testing"

func Test_quickSelect(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		arr  []int
		k    int
		want int
	}{
		{"test1", []int{4, 1, 2, 5, 3}, 0, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := quickSelect(tt.arr, tt.k)
			if got != tt.want {
				t.Errorf("quickSelect() = %v, want %v", got, tt.want)
			}
		})
	}
}
