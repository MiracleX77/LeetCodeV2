package programingskill

import (
	"fmt"
)

func S860(bills []int) bool {
	bills_dict := make(map[int]int)
	bills_dict[5] = 0
	bills_dict[10] = 0
	bills_dict[20] = 0
	for _, i := range bills {
		switch bill := i; bill {
		case 5:
			bills_dict[5] = bills_dict[5] + 1
		case 10:
			if bills_dict[5] == 0 {
				return false
			}
			bills_dict[5] = bills_dict[5] - 1
			bills_dict[10] = bills_dict[10] + 1
		case 20:
			if bills_dict[10] == 0 || bills_dict[5] == 0 {
				if bills_dict[5] < 3 {
					return false
				} else {
					bills_dict[5] = bills_dict[5] - 3
					bills_dict[20] = bills_dict[20] + 1
				}
			} else {
				bills_dict[5] = bills_dict[5] - 1
				bills_dict[10] = bills_dict[10] - 1
				bills_dict[20] = bills_dict[20] + 1
			}

		}
		fmt.Println(bills_dict)
	}
	fmt.Println(bills_dict)
	return true
}
