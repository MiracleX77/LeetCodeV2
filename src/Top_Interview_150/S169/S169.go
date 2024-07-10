package main

import "fmt"

func main() {
	nums1 := []int{
		2, 2, 1, 1, 1, 2, 2,
	}
	fmt.Println(majorityElement(nums1))
}

func majorityElement(nums []int) int {
	freqMap := make(map[int]int)

	// นับความถี่ของตัวเลขในอาร์เรย์
	for _, num := range nums {
		freqMap[num]++
	}
	fmt.Println(freqMap[2])
	return 0
}
