package main

import (
	"fmt"
	programingskill "mix/ProgramingSkill"
	//feb2024 "mix/Feb_2024"
)

func main() {
	// moves := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }
	// programingskill.S1523(3, 7)
	bills := []int{
		5, 5, 10, 10, 20,
	}
	result := programingskill.S860(bills)
	fmt.Println(result)
}
