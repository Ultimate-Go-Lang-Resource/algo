// Problem:
// Given a list of integer where every element appears even number of time
// except for one, find that unique one in O(1) space.
//
// Example:
// Input: []int{1, 1, 2, 4, 2, 1, 4, 6, 1}
// Return: 6, because 6 appears 1 time only.
//
// Solution:
// Bitwise XOR all integers in the list.
// Since XOR-ing a number with itself is zero, we could cancel out all the even
// duplicates.
// That leaves us the one that appears one time only.
//
// Cost:
// O(n) time, O(1) space.

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestFindUniqueID(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{6}, 6},
		{[]int{6, 2, 2}, 6},
		{[]int{6, 2, 2, 4, 4, 1, 1, 1, 1}, 6},
		{[]int{1, 1, 2, 4, 2, 1, 4, 6, 1}, 6},
		{[]int{4, 1, 2, 1, 6, 1, 4, 2, 1}, 6},
	}

	for _, tt := range tests {
		result := findUniqueID(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

func findUniqueID(list []int) int {
	uniqueID := 0

	for _, id := range list {
		uniqueID ^= id
	}

	return uniqueID
}
