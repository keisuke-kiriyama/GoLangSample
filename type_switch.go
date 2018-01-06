package main

import "fmt"

func main() {
	showType(nil)
	showType(12345)
	showType("abcdef")
	showType(3.14)
}

func showType(x interface{}) {
	switch val := x.(type) {
	case nil:
		fmt.Println("nil")
	case int, int32, int64:
		fmt.Println("整数", val)
	case string:
		fmt.Println("文字列", val)
	default:
		fmt.Println("不明", val)
	}
}
