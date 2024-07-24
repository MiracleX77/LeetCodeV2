package main

import (
	"fmt"
)

func main() {

	nums := []int{
		1, 2, 3, 4,
	}
	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	product := 1
	for i, num := range nums {
		res[i] = product
		product *= num
	}

	sufProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= sufProduct
		sufProduct *= nums[i]
	}

	return res
}
