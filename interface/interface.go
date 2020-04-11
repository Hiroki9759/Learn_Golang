package main

import "fmt"

//インターフェースとは中身(ロジック)を持たない関数のこと

//演算インターフェース型

type Calculator interface {
	//関数の定義
	Calculate(a int, b int) int
}

//足し算型

type Add struct {
}

//足し算型にCalculatorインターフェースのCalcurate関数を実装(具体化)

func (x Add) Calculate(a int, b int) int {
	return a + b
}

//引き算型
type Sub struct {
}

//引き算型にCalculatorインターフェースのCalcurate関数を実装(具体化)
func (x Sub) Calculate(a int, b int) int {
	return a - b
}

//掛け算型
type Mul struct {
}

//掛け算型にCalculatorインターフェースのCalcurate関数を実装(具体化)
func (x Mul) Calculate(a int, b int) int {
	return a * b
}

//割り算型
type Div struct {
}

//割り算型にCalculatorインターフェースのCalcurate関数を実装(具体化)
func (x Div) Calculate(a int, b int) int {
	return a / b
}

type Encoder interface {
	Encode(string) string
}
type Decoder interface {
	Decode(string) string
}
type Converter interface {
	//インターフェースの埋め込み
	Encoder
	Decoder
}

func main() {
	//実装した型の変数を宣言
	var add Add
	var sub Sub
	var mul Mul
	var div Div

	//Calculatorインターフェース型の変数を宣言
	var cal Calculator

	//add型の値を代入
	cal = add
	//インターフェース経由でメソッドを取り出す
	fmt.Println("和", cal.Calculate(1, 2))
	//sub型の値を代入
	cal = sub
	//インターフェース経由でメソッドを取り出す
	fmt.Println("差", cal.Calculate(1, 2))
	//mul型の値を代入
	cal = mul
	//インターフェース経由でメソッドを取り出す
	fmt.Println("積", cal.Calculate(1, 2))
	//div型の値を代入
	cal = div
	//インターフェース経由でメソッドを取り出す
	fmt.Println("商", cal.Calculate(1, 2))

	//空インターフェースにstring型の値を格納
	var i interface{} = "test"
	//型アサーションを使いstring型に変換
	//変換できない時はランタイムパニックを起こす
	var s string = i.(string)
	fmt.Println(s)
	//起こさないために
	s1, ok := i.(string)
	fmt.Println(s1, ok)
	s2, ok := i.(interface {
		dummy()
	})
	fmt.Println(s2, ok)
	showType(nil)
	showType(12345)
	showType(3.14)
	showType("asdfg")
}

//型判別関数 型switch文を使う
func showType(x interface{}) {
	switch val := x.(type) {
	case nil:
		fmt.Println("nil")
	case int, int16, int32, int64, int8:
		fmt.Println("int", val)
	case float32, float64:
		fmt.Println("float", val)
	case string:
		fmt.Println("文字列", val)
	default:
		fmt.Println("不明", val)
	}
}
