package main

import "fmt"

func main() {
LBL1:
	for i := 0; i < 5; i++ {
		switch {
		case i == 3:
			break LBL1
		default:
			fmt.Println(i)

		}
	}
LBL2:
	for i := 0; i < 3; i++ {

		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Println("   ", j)
			if i == 1 || j == 1 {
				continue LBL2
			}
		}
	}
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			goto EVEN
		}
		fmt.Println("odd", i)
	EVEN:
	}
}
