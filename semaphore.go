package main

import (
	"fmt"
	"math/rand"
	"time"
)

const goroutines = 10
const maxProcesses = 3

func main() {
	semaphore := make(chan int, maxProcesses)
	notify := make(chan int, goroutines)

	for i := 0; i < goroutines; i++ {
		go func(no int, semaphore chan int, norify chan<- int) {
			semaphore <- 0 //自分の番が来るのをまつ
			time.Sleep(
				time.Duration(rand.Int63n(3)) * time.Second)
			fmt.Println("処理完了", no)
			<-semaphore  //待機中のゴルーチンに処理を明け渡す
			notify <- no //処理完了を通知
		}(i, semaphore, notify)
	}

	for i := 0; i < goroutines; i++ {
		<-notify
	}

	fmt.Println("全て完了")
}
