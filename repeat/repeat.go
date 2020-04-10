package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	k := 0
	for k < 5 {
		fmt.Println("while", k)
		k++
	}
	l := 0
	for {
		fmt.Println("無限ループ")
		l++
		if l == 5 {
			break
		}
	}
	arr := [...]int{0, 1, 2, 3, 4}
	for i := range arr {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
