package main

import (
	"fmt"
	"sync"
)

var sy = sync.WaitGroup{} // 开启一个等待组来阻塞程序不会终止，等所有协程执行完毕之后在关闭
func showMsg(str string) {
	sy.Done()
	fmt.Printf("str:%v\n", str)
}

func main() {

	for i := 0; i < 5; i++ {
		sy.Add(1)
		go showMsg("java") // 开启一个携程来执行 showMsg方法
	}
	sy.Wait() // 全部协程执行完毕之后才继续执行
	fmt.Println("end...\n")
}
