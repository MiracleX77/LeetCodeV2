package main

import "fmt"

func main() {

	nums1 := []int{
		0,
	}
	nums2 := []int{
		1, 3,
	}
	m := 0
	n := 2
	fmt.Println(merge(nums1, m, nums2, n))
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	nums1 = nums1[:m]
	nums2 = nums2[:n]

	if m == 0 {
		nums1 = nums2
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if nums2[i] < nums1[j] {
				count := nums1[j]
				nums1[j] = nums2[i]
				nums1 = append(nums1, count)
				break
			}
			if j == m-1 {
				nums1 = append(nums1, nums2[i])
			}
		}
	}
	return nums1
}
