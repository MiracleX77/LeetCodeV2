package programingskill

import (
	"sort"
)

func S2966(nums []int, k int) [][]int {
	result := [][]int{}
	num := []int{}
	sort.Ints(nums)
	for i, v := range nums {
		num = append(num, v)
		if i%3 == 2 {
			for i := 1; i < 3; i++ {
				if num[i]-num[0] > k {
					result = [][]int{}
					return result
				}
			}
			result = append(result, num)
			num = []int{}
		}
	}
	return result
}
