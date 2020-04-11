package main

import "fmt"

type Person struct {
	name string
	age  int
}

//構造体の埋め込み
type Employee struct {
	id int
	Person
}

func main() {
	//フィールドを個別に初期化
	var p1 Person
	p1.name = "Kondo"
	p1.age = 27

	//構造体リテラルで初期化
	p2 := Person{age: 31, name: "Goto"}
	p3 := Person{"Ito", 45}
	//ポインタも構造体リテラルで作成可能
	p4 := &Person{"Saito", 67}
	fmt.Println(p1, p2, p3, p4)
	//初期化の注意　最後のデータと}が同じ行&&最後のデータの後ろに,が来て改行}
	e := Employee{1, Person{"Iwana", 23}}
	fmt.Println(e)
}
