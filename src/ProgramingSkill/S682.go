package programingskill

import (
	"fmt"
	"strconv"
)

func S682(ope []string) int {
	result := []int{}
	for _, v := range ope {
		switch v {
		case "+":
			intVar_1 := result[len(result)-1]
			intVar_2 := result[len(result)-2]
			result = append(result, (intVar_1 + intVar_2))
		case "D":
			intVar_1 := result[len(result)-1]
			result = append(result, (intVar_1 * 2))
		case "C":
			result = result[:len(result)-1]
		default:
			intVar, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("Error converting string to int: %v\n", err)
				continue
			}
			result = append(result, intVar)
		}

	}
	total := 0
	for _, i := range result {
		total += i
	}
	return total
}
