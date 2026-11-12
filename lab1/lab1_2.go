package main

func calcBoolLogic(x float64) (float64, bool) {
	var y float64

	if (x >= 0 && x <= 8) || (x >= 16) {
		y = -3*x*x/5 + 9
	} else if x < -1 {
		y = 15*x - 2
	} else {
		return 0, false
	}
	return y, true
}
