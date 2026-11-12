package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("n | вкладені цикли | динамічне програмування")

	for n := 1; n <= 3; n++ {
		s1, _ := calcNested(n)
		s2, _ := calcDP(n)
		fmt.Printf("%d | %.7f | %.7f\n", n, s1, s2)
	}

	f, err := os.Create("ops.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	fmt.Fprintln(f, "n,ops_nested,ops_dp")
	for n := 1; n <= 20; n++ {
		_, o1 := calcNested(n)
		_, o2 := calcDP(n)
		fmt.Fprintf(f, "%d,%d,%d\n", n, o1, o2)
	}
	fmt.Println("\nops.csv записано (n = 1..20)")

}
