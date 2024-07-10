package main

import (
	"fmt"
)

func main() {
	nums1 := []int{
		1, 2, 3, 4, 5, 6, 7,
	}
	k := 3
	fmt.Println(rotate(nums1, k))
}

func rotate(nums []int, k int) []int {
	n := len(nums)
	k = k % n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
	return nums

}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
