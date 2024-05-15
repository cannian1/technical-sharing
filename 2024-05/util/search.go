package util

import (
	"cmp"
)

// BinarySearch 二分查找
func BinarySearch[S ~[]T, T cmp.Ordered](nums S, target T) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := ((right - left) >> 1) + left
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
