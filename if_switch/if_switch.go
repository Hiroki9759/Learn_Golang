package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			println(i, "は偶数")
		} else {
			println(i, "は奇数")
		}
	}
	i := 0
	for i < 10 {
		i++
		switch i {
		case 0:
			fmt.Println("0")
		case 1, 2:
			fmt.Println("1または2")
		default:
			fmt.Println("その他")
		}
	}
	for day := time.Sunday; day <= time.Saturday; day++ {
		switch day {
		case time.Sunday:
			fallthrough
		case time.Saturday:
			fmt.Println("Holiday")
		default:
			fmt.Println("平日")
		}

	}
	for i := -2; i <= 2; i++ {
		switch true {
		case i > 0:
			fmt.Println("+")
		case i < 0:
			fmt.Println("-")
		default:
			fmt.Println("0")
		}
	}

}
