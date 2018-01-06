package main

import (
	"fmt"
	"math/rand"
	"time"
)

const goroutines = 10

func main() {
	c := make(chan int)

	for i := 0; i < goroutines; i++ {
		go func(s chan<- int) {
			time.Sleep(
				time.Duration(rand.Int63n(10)) * time.Second)

			fmt.Println("処理完了")

			s <- 0 //処理完了を伝える
		}(c)
	}

	for i := 0; i < goroutines; i++ {
		<-c //channelのバッファが0だから受信があるまで止まる。それによって同期
	}

	fmt.Println("全て完了")
}
