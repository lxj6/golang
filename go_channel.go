package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ch = make(chan int) // 创建一个无缓冲通道

func testchan() {
	val := rand.Intn(10)        // 创建一个随机数
	fmt.Printf("val:%v\n", val) // 打印
	time.Sleep(time.Second * 3) // 延迟三秒在放进通道
	ch <- val
}

func main() {
	defer close(ch)           // 关闭通道
	go testchan()             // 启用协程执行
	num := <-ch               // 取出通道内的值
	fmt.Printf("num:%v", num) // 打印
}
