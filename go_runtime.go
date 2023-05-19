package main

import (
	"fmt"
	"runtime"
)

func show(str string) {
	for i := 0; i < 2; i++ {
		fmt.Printf("str:%v\n", str)
	}
}

func main() {
	go show("java") // 子协程
	// 主协程
	for i := 0; i < 2; i++ {
		runtime.Gosched() // 让出cpu执行权利，优先执行子协程
		fmt.Println("golang..")
	}
}
