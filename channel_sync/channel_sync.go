package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const goroutines = 10
const maxProcesses = 3

func main() {
	semaphore := make(chan int, maxProcesses)
	notify := make(chan int, goroutines)
	//値の教養に使用するチャネルの作成
	counter := make(chan int, 1)
	for i := 0; i < goroutines; i++ {
		go func(counter chan int) {
			value := <-counter
			value++
			fmt.Println("counter:", value)
			if value == goroutines {
				os.Exit(0)
			}
			counter <- value
		}(counter)
	}

	counter <- 0

	for i := 0; i < goroutines; i++ {
		go func(no int, semaphore chan int, notify chan<- int) {
			semaphore <- 0
			time.Sleep(
				time.Duration(rand.Int63n(3)) * time.Second)
			fmt.Println("処理完了", no)

			<-semaphore
			notify <- no
		}(i, semaphore, notify)
	}
	for i := 0; i < goroutines; i++ {
		<-notify
	}

	c := make(chan int)
	for i := 0; i < goroutines; i++ {
		go func(s chan<- int) {
			time.Sleep(
				time.Duration(rand.Int63n(10)) * time.Second)
			fmt.Println("処理完了")
			s <- 0
		}(c)
	}
	for i := 0; i < goroutines; i++ {
		<-c
	}
	//全て完了
	fmt.Println("全て完了")
}

//セマフォ(並列処理においてリソースへの同時アクセスを抑制する仕組み)の再現もできる
