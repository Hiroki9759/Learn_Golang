package main

import "fmt"

func main() {

	var b bool

	b = true
	fmt.Println(b)
	b = false
	fmt.Println(b)
	b = true || false
	fmt.Println(b)

	var i int = 12345
	var i64 int64 = int64(i)
	var u uint32 = uint32(i)
	var f float32 = float32(i)
	var s string = string(i)
	var by []byte = []byte("abc")
	fmt.Println(i64, u, f, s, by)

	var en string = "golang"
	var str string = "あ"
	str = str + "い"
	str += "う"
	fmt.Println(str)
	fmt.Println(en, "len:", len(en))
	fmt.Println(str, "len:", len(str))
	type myInteger int
	var j myInteger = 1223
	j++
	fmt.Println(j)
	type myStruct struct {
		a int
		b int
	}
	k := 124
	ii, st := i, "hoge"
	println(k, ii, st)

	const G float32 = 9.80665
	const G2 = G * 2
	const π float64 = 3.14159265
	const π2 float64 = π * 2
	const (
		X int = 1
		Y
		Z
	)
	println(X, Y, Z, π2)

	//iotaは列挙時に便利
	const (
		ZERO  = iota
		ONE   = iota
		TWO   = iota
		THREE = iota
		FOUR
	)
	println(ZERO, ONE, TWO, THREE, FOUR)
	const (
		bit0 = 1 << iota
		bit1
		bit2
		bit3
	)
	println(bit0, bit1, bit2, bit3)
}
