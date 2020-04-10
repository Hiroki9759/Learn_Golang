package main

import "fmt"

func main() {
	var ptr *int
	var i int = 12345
	ptr = &i
	fmt.Printf("%#v\n", "アドレス")
	fmt.Printf("%#v\n", &i)
	fmt.Printf("%#v\n", ptr)
	fmt.Printf("%#v\n", "変数")
	fmt.Printf("%#v\n", i)
	fmt.Printf("%#v\n", *ptr)
	*ptr = 999
	fmt.Printf("%#v\n", i)

	a, b := 1, 1
	double(a, &b)
	fmt.Println("値渡し", a)
	fmt.Println("ポインタ渡し", b)
	//[new組み込み関数によるメモリの割り当て]
	var c *int = new(int)
	var s *string = new(string)
	*c, *s = 1, "asd"
	fmt.Println(*c, *s)

}

func double(x int, y *int) {
	x = x * 2
	*y = *y * 2
}
