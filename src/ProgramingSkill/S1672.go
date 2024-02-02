package programingskill

func S1672(accounts [][]int) int {
	result := 0
	for _, v := range accounts {
		sum := 0
		for _, i := range v {
			sum += i
		}
		if sum >= result {
			result = sum
		}
	}
	return result
}
