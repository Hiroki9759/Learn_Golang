package main

import "fmt"

func main() {
	//マップ：キーと値の組み合わせ
	//マップの作成
	capitals := make(map[string]string)
	capitals["日本"] = "東京"
	capitals["アメリカ"] = "ワシントンDC"
	capitals["中国"] = "北京"
	for country, capital := range capitals {
		fmt.Println(country, capital)
	}
	//キーの存在確認
	_, ok := capitals["イギリス"]
	if ok {
		fmt.Println("登録済み")
	} else {
		fmt.Println("未登録")
	}
	//マップの初期化　マップリテラルを使用
	capitals1 := map[string]string{
		"日本":   "東京",
		"アメリカ": "ワシントンDC",
		"中国":   "北京",
	}
	fmt.Println(capitals1)

}
