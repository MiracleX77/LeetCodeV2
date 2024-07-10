package main

import "fmt"

func main() {

	nums := []int{
		2, 3, 1, 1, 4,
	}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	count_jump := nums[0]
	for i := 1; i < len(nums); i++ {
		if count_jump == 0 {
			return false
		}
		count_jump = max(count_jump-1, nums[i])
	}
	return true
}
