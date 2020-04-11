package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 10)
	//送信専用チャネル
	go func(s chan<- int) {
		//チャンネルへ０〜４の値を送信
		for i := 0; i < 5; i++ {
			s <- i
		}
		close(s)
	}(c)

	//受信ループ
	for {
		value, ok := <-c
		//キャパシティを確認
		fmt.Println("cap:", cap(c))
		//現在の要素数を確認
		fmt.Println("len", len(c))
		if !ok {
			break
		}
		fmt.Println(value)
	}

}
