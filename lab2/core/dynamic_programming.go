package main

import "math"

func calcDP(n int) (float64, int) {
	opsCount := 0
	summary := 0.0
	numerator := 1.0
	for i := 1; i <= n; i++ {
		cosI := math.Cos(float64(i))
		numerator *= float64(i) - cosI
		summary += numerator / (float64(i) + cosI)
		opsCount += 9
		opsCount++
	}
	return summary, opsCount
}
