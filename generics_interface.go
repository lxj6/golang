package main

import "fmt"

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

// 泛型切片的定义
type List[T any] []T

// 泛型集合的定义。声明了两个泛型，分别为k和v
type MapT[k comparable, v any] map[k]v

// 泛型通道的定义
type Chan[T any] chan T

// 基本接口的定义(可用于变量的定义)
type ToString interface {
	string() string
}

//var s toString

// 泛型接口的定义
type GetKey[T comparable] interface {
	any
	Get() T
}

func (u user) String() string {
	return fmt.Sprintf("ID: %d,Name: %s,Age: %d", u.ID, u.Name, u.Age)
}

func (u user) Get() int64 {
	return u.ID
}

func (a address) String() string {
	return fmt.Sprintf("ID: %d,Province: %s,City: %s", a.ID, a.Province, a.City)
}

func (a address) Get() int {
	return a.ID
}

// 列表转集合
func listToMap[k comparable, T GetKey[k]](list []T) map[k]T {
	mp := make(MapT[k, T], len(list))
	for _, data := range list {
		mp[data.Get()] = data
	}
	return mp
}

func interfaceCase() {
	userList := []GetKey[int64]{
		user{1, "a", 19},
		user{2, "b", 20},
	}
	addrList := []GetKey[int]{
		address{1, "北京", "北京"},
		address{2, "天津", "天津"},
	}
	userMp := listToMap(userList)
	fmt.Println(userMp)
	addrMp := listToMap(addrList)
	fmt.Println(addrMp)
}

func main() {
	interfaceCase()
}
