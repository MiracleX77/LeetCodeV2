package main

import (
	"fmt"
)

func main() {
	nums1 := []int{
		5, 1, 2, 3, 4,
	}
	nums2 := []int{
		4, 4, 1, 5, 1,
	}

	fmt.Println(canCompleteCircuit(nums1, nums2))
}

func canCompleteCircuit(gas []int, cost []int) int {

	max_profit := -1
	profit := make([]int, len(gas))
	index := 0
	for i := 0; i < len(gas); i++ {
		profit[i] = gas[i] - cost[i]
		if profit[i] > max_profit {
			index = i
			max_profit = profit[i]
		}
	}
	sum := 0
	for n := 0; n < len(gas); n++ {
		sum += profit[n]
	}
	if sum < 0 {
		return -1
	}
	fmt.Println(profit)
	return index
}
