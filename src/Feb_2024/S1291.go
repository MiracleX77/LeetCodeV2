package feb2024

import (
	"fmt"
	"strconv"
)

func S1291(low int, high int) []int {
	str := "123456789"
	result := []int{}
	last_int := low
	low_str := strconv.Itoa(low)
	low_first_int, _ := strconv.Atoi(string(low_str[0]))
	index_str := len(low_str)
	//fmt.Print(low_first_int)
	for last_int < high {
		last_str := ""
		for i := 0; i < index_str; i++ {
			last_str += string(str[low_first_int-1+i])
		}
		low_first_int++
		fmt.Println(last_str)
		last_int, _ = strconv.Atoi(last_str)
		result = append(result, last_int)
		fmt.Println(result)
	}
	//char := []rune(str)[0]
	//fmt.Print(string(char))
	return []int{}
}
