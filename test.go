package main

import "fmt"

type usb interface {
	read()
	write()
}

type mi struct {
}

func (m mi) read() {
	fmt.Println("mi read...\n")
}

func (m mi) write() {
	fmt.Println("mi write...\n")
}

func main() {
	var us usb = mi{}
	us.write()
	us.read()
}
