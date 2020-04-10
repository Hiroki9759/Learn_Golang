package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}
func calc(a int, b int) (int, int, int, float32) {
	return a + b, a - b, a * b, float32(a) / float32(b)
}
func calc2(a int, b int) (add2 int, sbb2 int, mul2 int, div2 float32) {
	add2 = a + b
	sbb2 = a - b
	mul2 = a * b
	div2 = float32(a) / float32(b)
	return
}
func f1() (int, string, float32) {
	return 0, "xy", 3.14
}
func f2(a int, b string, c interface{}) {
	fmt.Println(a, b, c)
}

func holiday(month int, days ...string) {
	fmt.Printf("%d月の祝日は%d日あります。\n", month, len(days))
	for _, day := range days {
		fmt.Println(day)
	}
	fmt.Println()
}
func main() {
	answer := plus(1, 2)
	add, sbb, mul, div := calc(1, 2)

	fmt.Println(answer, add, sbb, mul, div)
	f2(f1())
	holiday(1, "元旦", "成人の日")
	holiday(2, "建国記念の日")
	holiday(3, "春分の日")
	add2, sbb2, mul2, div2 := calc2(1, 2)
	fmt.Println(add2, sbb2, mul2, div2)
}
