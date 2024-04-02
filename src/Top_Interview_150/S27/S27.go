package main

import "fmt"

func main() {

	nums := []int{
		0, 1, 2, 2, 3, 0, 4, 2,
	}
	val := 2
	fmt.Println(removeElement(nums, val))
}

func removeElement(nums []int, val int) int {
	count := 0
	last_index := len(nums) - 1
	n := 0
	for n < len(nums) {
		if nums[n] == val {
			for i := n; i < len(nums)-1; i++ {
				nums[i] = nums[i+1]
			}
			count++
			nums[last_index] = -1
		} else {
			n++
		}
	}
	return len(nums) - count
}
