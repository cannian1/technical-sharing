package util

import "fmt"

func ExampleBinarySearch() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 8

	index := BinarySearch(nums, target)
	fmt.Println(index)
	// Output:
	// 7
}
