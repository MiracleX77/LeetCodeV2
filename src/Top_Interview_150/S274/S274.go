package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{
		3, 1, 1,
	}
	fmt.Println(hIndex(nums))
}

func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	for i, cucitation := range citations {
		if cucitation < i+1 {
			return i
		}
	}
	return len(citations)
}
