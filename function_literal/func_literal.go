package main

import "fmt"

func init() {
	fmt.Println("初期化")
}
func main() {

	val := 123
	//記述と呼び出しを同時に行う
	func(i int) {
		fmt.Printf("%#v\n", i*val)
	}(10)

	//関数リテラルを変数の代入
	f := func(s string) {
		fmt.Printf("%#v\n", s)
	}
	f("hoge")
	//関数型の変数宣言
	var f1 func(int, int) int
	//関数型の変数に関数リテラルの値を代入
	f1 = func(a int, b int) int {
		return a + b
	}
	//関数型の変数経由で関数呼びだす
	fmt.Println(f1(1, 2))
	//関数型の変数に関数のあたいを代入
	f1 = multiply
	//関数型の変数経由で関数呼びだす
	fmt.Println(f1(1, 2))

}
func multiply(x int, y int) int {
	return x * y
}

/*関数リテラル：匿名の関数　他の関数内に記述する
関数リテラルの外側の変数に対してアクセスができる
このような関数のことを「クロージャ」と呼ぶらしい
呼び出しと定義を一気にするか、関数リテラルを一旦変数に代入してからその変数を経由してよび出しを行う*/
