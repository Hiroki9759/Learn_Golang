package main

import (
	"fmt"
	"os"
)

type myType int

func (value myType) println() {
	fmt.Println(value)
}

//レシーバが非ポインタ
func (value myType) setByValue(newValue myType) {
	value = newValue
}

//レシーバがポインタ
func (value *myType) setByPointer(newValue myType) {
	*value = newValue
}

//値を加算するメソッド
func (value *myType) add(increment myType) myType {
	*value += increment
	return *value
}

func delay() {
	fmt.Println("delay called")
}
func main() {
	file, err := os.OpenFile("text.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString("アイウエオ")
	var z myType = 1234
	z.println()
	defer delay()
	var x myType = 0
	x.setByValue(1)
	fmt.Println(x)
	x.setByPointer(1)
	fmt.Println(x)
	var i myType
	//iに対しaddメソッドで３を加算
	fmt.Println(i.add(3))
	//メソッド値を取得　型は「func(myType) myType」
	mv := i.add
	fmt.Println(mv(3))

}
