package main

import "fmt"

//埋め込まれる側の構造体
type embedded struct {
	i int
}

func (x embedded) doSomething() {
	fmt.Println("test.doSomething()")
}

type test struct {
	embedded
}

func main() {
	var x struct {
		i1, i2 int     //int Field
		f      float32 //float Field
		s      string  //string Field
	}
	x.i1 = 1
	x.i2 = 2
	x.s = "go"
	x.f = 3.14
	fmt.Println(x)
	//ブランクフィールド
	// struct {
	// 	b1	byte
	// 	_	byte
	// 	b2	[5]byte
	// 	_	byte
	// }
	var y test
	y.doSomething()
	fmt.Println(y.i)
}
