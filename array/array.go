package main

import "fmt"

func main() {
	var array1 [1]byte
	var array2 [5]*int
	var array3 [8][3]int64
	var array4 [2]struct{ x, y int }

	fmt.Println(len(array1))
	fmt.Println(len(array2))
	fmt.Println(len(array3), len(array3[0]))
	fmt.Println(len(array4))
	var date [7]string
	date[0] = "Sunday"
	date[1] = "Monday"
	date[2] = "Tuesday"
	date[3] = "Wednesday"
	date[4] = "Thursday"
	date[5] = "Friday"
	date[6] = "Saturday"
	i := 0
	for i < len(date) {
		fmt.Println(date[i], " ")
		i++
	}
	fmt.Println()

	for _, value := range date {
		fmt.Println(value, " ")
	}
	fmt.Println()

	//配列の初期化
	array6 := [5]float32{}
	array7 := [6]int{1, 2, 3, 4}
	array8 := [...]string{"One", "Two", "Three"}
	fmt.Println(array6, array7, array8)

	x := [5]string{"a", "b", "c", "d", "e"}
	//スライス型の変数を宣言
	var s1 []string
	//配列全体をスライス
	s1 = x[:]
	fmt.Println(s1)
	s2 := x[1:4]
	fmt.Println(s2)
	s3 := x[:3]
	fmt.Println(s3)
	s4 := x[3:]
	fmt.Println(s4)
	//スライスは参照型(マップ,スライス,チャネルの3つある)
	//配列の場合関数で値の受け渡しをすると、受け渡し先で要素が全てコピーされる
	values := [...]int{0, 1, 2, 3, 4}
	double(values[:])
	fmt.Println(values)

	var w string = "abcde"[1:4]
	var y string = "あいうえお"[3:9] //平仮名は３byteなので”いう”が取り出される
	var z string = "あいうえお"[1:8] //変なスライスをすると文字化けを起こす
	fmt.Println(w)
	fmt.Println(y)
	fmt.Println(z)
	//要素の追加
	slice1 := []int{0, 1, 2, 3, 4}
	slice2 := append(slice1, 5, 6, 7, 8, 9)
	//要素のコピー
	fmt.Println(slice2)
	count := copy(slice1[2:], slice2)
	fmt.Println("count:", count)
	fmt.Println(slice1)
	w1 := []string{"a", "b", "c", "d", "e"}
	w2 := append(w1, "f", "g")
	w3 := append(w1, w2...)
	fmt.Println(w2, w3)

	//スライスの初期化
	slice3 := make([]int, 10, 20) //長さ10キャパシティ20のスライスを作成
	fmt.Println(slice3)
	fmt.Println("len", len(slice3))
	fmt.Println("capacity", cap(slice3))
	//リテラルによるスライスの初期化
	s := []int{1, 2, 3, 4}
	fmt.Println(s)
}
func double(values []int) {
	for i := 0; i < len(values); i++ {
		values[i] *= 2
	}
}
