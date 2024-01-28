package programingskill

import (
	"fmt"
	"math"
	"strconv"
)

func S1491(salary []int) float64 {
	sum := 0.0
	max := math.Inf(-1)
	min := math.Inf(1)
	for i := range salary {
		v := float64(salary[i])
		if v >= max {
			max = v
		}
		if v <= min {
			min = v
		}
	}
	for i := range salary {
		v := float64(salary[i])
		if v != max && v != min {
			sum += v
		}
	}

	resStr := fmt.Sprintf("%.5f", sum/float64(len(salary)-2))
	res, err := strconv.ParseFloat(resStr, 64)
	if err != nil {
		return 0.0
	}
	fmt.Print(res)
	return res
}
