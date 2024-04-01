package programingskill

func S1232(coordinates [][]int) bool {
	if len(coordinates) < 2 {
		return false
	}
	diff_x := float64(coordinates[1][0] - coordinates[0][0])
	diff_y := float64(coordinates[1][1] - coordinates[0][1])
	for i := 1; i < len(coordinates); i++ {
		diff_x_1 := float64(coordinates[i][0] - coordinates[i-1][0])
		diff_y_1 := float64(coordinates[i][1] - coordinates[i-1][1])
		if diff_x*diff_y_1 != diff_x_1*diff_y {
			return false
		}
	}
	return true
}
