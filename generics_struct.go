package main

import "fmt"

// 泛型结构体
type MyStruct[T interface{ *int | *string }] struct {
	Name string
	Date T
}

// 泛型receiver
func (my MyStruct[T]) GetDate() T {
	return my.Date
}

func main() {
	data := 18
	mystruct := MyStruct[*int]{
		Name: "s",
		Date: &data,
	}
	data1 := mystruct.GetDate()
	fmt.Println(*data1)
}
