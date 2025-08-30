package main

import (
	"fmt"
	"time"
)

func main() {
	var m = [...]int{1, 2, 3, 4, 5}
	for i, v := range m {
		go func(i, v int) {
			time.Sleep(time.Second * 3)
			fmt.Println(i, v)
		}(i, v) // 将i, v的值作为参数传递给匿名函数
	}
	time.Sleep(time.Second * 10)

	for i, v := range m {
		go func() {
			time.Sleep(time.Second * 3)
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second * 10)
}
