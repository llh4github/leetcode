package array

import (
	"slices"
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums   []int
		target int
		want   []int
	}{
		{"test1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			slices.Sort(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		{"test1", []int{1, 1, 2}, 2},
		{"test2", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.nums)
			t.Log(tt.nums)
			if got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeElement(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		val  int
		want int
	}{
		{"test1", []int{3, 2, 2, 3}, 3, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.nums, tt.val)
			if got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_plusOne(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		digits []int
		want   []int
	}{
		{"test1", []int{1, 2, 3}, []int{1, 2, 4}},
		{"test2", []int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plusOne(tt.digits)
			if !slices.Equal(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
