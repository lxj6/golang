package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

type user struct {
	ID   int64
	Name string
	Age  uint8
}

type address struct {
	ID       int
	Province string
	City     string
}

func main() {
	TTypte()
	TTypte1()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}

// 集合转列表
func mapToList[K comparable, T any](mp map[K]T) []T {
	list := make([]T, len(mp))
	var i int
	for _, value := range mp {
		list[i] = value
		i++
	}
	return list
}

func myPrintln[T any](ch chan T) {
	for data := range ch {
		fmt.Println(data)
	}
}

func TTypte() {
	userMp := make(map[int64]user, 0)
	userMp[1] = user{1, "a", 18}
	userMp[2] = user{2, "b", 10}
	userList := mapToList(userMp)

	ch := make(chan user)
	go myPrintln(ch)
	for _, u := range userList {
		ch <- u
	}

	addrMp := make(map[int]address, 0)
	addrMp[1] = address{1, "北京", "北京"}
	addrMp[2] = address{2, "天津", "天津"}
	addrList := mapToList(addrMp)
	ad := make(chan address)
	go myPrintln(ad)
	for _, a := range addrList {
		ad <- a
	}
}

// 泛型切片的定义
type List[T any] []T

// 泛型集合的定义。声明了两个泛型，分别为k和v
type MapT[k comparable, v any] map[k]v

// 泛型通道的定义
type Chan[T any] chan T

// 使用泛型的集合转泛型切片赋值给泛型通道并打印
func TTypte1() {
	userMp := make(MapT[int64, user], 0)
	userMp[1] = user{1, "a", 18}
	userMp[2] = user{2, "b", 10}
	var userList List[user]
	userList = mapToList(userMp)

	ch := make(Chan[user])
	go myPrintln(ch)
	for _, u := range userList {
		ch <- u
	}

	addrMp := make(MapT[int, address], 0)
	addrMp[1] = address{1, "北京", "北京"}
	addrMp[2] = address{2, "天津", "天津"}
	var addrList List[address]
	addrList = mapToList(addrMp)
	ad := make(Chan[address])
	go myPrintln(ad)
	for _, a := range addrList {
		ad <- a
	}
}
