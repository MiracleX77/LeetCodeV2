package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{
		2, 2, 1, 1, 1, 2, 2,
	}
	fmt.Println(majorityElement(nums1))
}

func majorityElement(nums []int) int {
	sort.Ints(nums)
	half := float32(len(nums)) / 2.0
	count := 0
	corrent_num := -1

	for _, num := range nums {
		if num != corrent_num {
			corrent_num = num
			count = 1
		} else {
			count++
		}
		if float32(count) >= half {
			return corrent_num
		}
	}
	return corrent_num
}
