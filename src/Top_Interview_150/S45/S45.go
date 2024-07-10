package main

import "fmt"

func main() {

	nums := []int{
		2, 3, 1, 1, 4,
	}
	fmt.Println(jump(nums))
}

func jump(nums []int) int {

	count := 0
	max := len(nums) - 1

	for max != 0 {
		for i := 0; i < max; i++ {
			if i+nums[i] >= max {
				max = i
				count++
				break
			}
		}
	}
	return count
}
