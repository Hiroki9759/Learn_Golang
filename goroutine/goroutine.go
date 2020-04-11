package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	fmt.Println("main:開始")
	fmt.Println("testを通常の関数として実行")
	test()
	fmt.Println("testをゴルーチンとして実行")
	go test()
	time.Sleep(3 * time.Second)
	fmt.Println("main:終了")
	go func() {
		fmt.Println("Goroutine")
		os.Exit(0)
	}()
	for {
		runtime.Gosched()
	}
}
func test() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

/*ゴルーチンは並列処理の仕組みのこと
ファイルの読み込みとキー入力などの複数の操作を￥独立して行うことで処理の複雑化を怪回避できる
ゴルーチンは独立して動作する実行単位でありGo言語のランタイムによって管理され、必要に応じてオペレーティングシステムが管理するスレッドに割り当てられて実行します。*/
/*ゴルーチンのスレッドは基本1つなので増やしたい時はruntime.Gosched関数を呼ぶ*/
