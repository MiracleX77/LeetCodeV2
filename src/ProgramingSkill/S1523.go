package programingskill

import "fmt"

func S1523(low int, high int) int {
	fmt.Print((high - 2 - low) / 2)
	return (high-2-low)/2 + (low % 2) + (high & 2)
}
