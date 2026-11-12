package main

import (
	"fmt"
	"time"
)

const (
	delay = 10 * time.Millisecond
	cols  = 80
	rows  = 24
)

func gotoxy(x, y int) {
	fmt.Printf("\033[%d;%dH", y, x)
}
func clearScreen() {
	fmt.Printf("\033[2J\033[?25l")
}

func returnCursor() {
	fmt.Printf("\033[?25h")
}

func main() {
	clearScreen()
	defer returnCursor()

	centerX := cols / 2

	for k := range centerX {
		leftX := centerX - k
		rightX := centerX + 1 + k

		for r := range rows {
			topDownY := 1 + r
			bottomUpY := rows - r

			var leftY, rightY int
			if k%2 == 0 {
				leftY = topDownY
				rightY = bottomUpY
			} else {
				leftY = bottomUpY
				rightY = topDownY
			}

			gotoxy(leftX, leftY)
			fmt.Print("*")
			time.Sleep(delay)

			gotoxy(rightX, rightY)
			fmt.Print("*")
			time.Sleep(delay)
		}
	}
	gotoxy(1, rows+1)
}
