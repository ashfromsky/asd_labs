package main

func calcSimpleComparisons(x float64) (float64, bool) {
	var y float64
	if x < -1 {
		y = 15*x - 2
	} else if x < 0 {
		return 0, false
	} else if x <= 8 {
		y = -3*x*x/5 + 9
	} else if x >= 16 {
		y = -3*x*x/5 + 9
	} else {
		return 0, false
	}
	return y, true
}
