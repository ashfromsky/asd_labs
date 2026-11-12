package main

import "math"

func calcNested(n int) (float64, int) {
	opsCount := 0
	summary := 0.0
	for i := 1; i <= n; i++ {
		numerator := 1.0
		opsCount++
		for j := 1; j <= i; j++ {
			numerator *= float64(j) - math.Cos(float64(j))
			opsCount += 4
			opsCount++
		}
		cosI := math.Cos(float64(i))
		summary += numerator / (float64(i) + cosI)
		opsCount += 6
		opsCount++
	}
	return summary, opsCount
}
