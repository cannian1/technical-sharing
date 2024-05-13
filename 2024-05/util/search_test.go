package util

import "testing"

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 5
	got := BinarySearch(nums, target)
	want := 4
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
