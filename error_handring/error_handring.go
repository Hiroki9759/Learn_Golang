package main

import (
	"fmt"
)

type MyError struct {
	message string //エラー詳細
}

//Errorメソッドの実装
func (err MyError) Error() string {
	return err.message
}

//16進数文字列をint型に変換
func hex2int(hex string) (val int, err error) {
	for _, r := range hex {
		val *= 16
		switch {
		case '0' <= r && r <= '9':
			val += int(r - '0')
		case 'a' <= r && r <= 'f':
			val += int(r-'a') + 10
		default:
			return 0, MyError{"不正な文字です。：" + string(r)}
		}
	}
	return
}

func main() {
	//fie open
	// file, err := os.Open("test,txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// }
	// file.Close()
	// fmt.Println("OK")

	val, err := hex2int("1")
	fmt.Println(val, err)
	val, err = hex2int("abcd")
	fmt.Println(val, err)
	val, err = hex2int("z")
	fmt.Println(val, err)

}
