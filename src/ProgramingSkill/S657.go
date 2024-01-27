package programingskill

func S657(moves string) bool {
	x, y := 0, 0
	for _, v := range moves {
		switch v {
		case 85:
			y += 1
		case 68:
			y -= 1
		case 82:
			x += 1
		case 76:
			x -= 1
		}
	}
	if x != 0 || y != 0 {
		return false
	}
	return true

}
