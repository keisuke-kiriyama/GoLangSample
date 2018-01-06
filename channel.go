package main

import "fmt"

func main() {
	c := make(chan int)

	go func(s chan<- int) {	//パラメータは送信専用チャネル
		for i := 0; i < 5; i++ {
			s <- i //送信
		}
		close(s)
	}(c)

	for {
		value, ok := <-c //受信 okはchanがcloseか

		if !ok {
			break
		}

		fmt.Println(value)
	}
}
