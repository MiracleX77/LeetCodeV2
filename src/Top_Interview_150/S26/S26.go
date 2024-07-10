package main

import "fmt"

func main() {

	nums := []int{
		0,0,1,1,1,2,2,3,3,4,
	}
	fmt.Println(removeElement(nums))
}

func removeElement(nums []int) int {
	left := 1
	for right := 1; right < len(nums); right++ {
		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
		fmt.Println(nums)
	}

	return left
}
