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
	if low > 123456789 {
		return result
	}
	for last_int < high {
		last_str := ""

		if low_first_int-1+index_str > 9 {
			index_str++
			low_first_int = 1
		}
		for i := 0; i < index_str; i++ {
			i_str := string(str[low_first_int-1+i])
			last_str += i_str
		}
		low_first_int++
		fmt.Println(last_str)
		last_int, _ = strconv.Atoi(last_str)
		if last_int > high {
			break
		} else {
			if last_int >= low {
				result = append(result, last_int)
			}
			if last_int == 123456789 {
				break
			}
			if last_int%10 == 9 {
				index_str++
				low_first_int = 1
			}
		}
		fmt.Println(result)
	}
	return result
}
