package main

import "fmt"

type player interface {
	playMusic()
}

type video interface {
	playVideo()
}

type mobile struct {
}

func (mo mobile) playMusic() {
	fmt.Println("playMusic....")
}
func (mo mobile) playVideo() {
	fmt.Println("playVideo...")
}

type peter interface {
	eat()
}

type dog struct {
}
type cat struct {
}

func (d dog) eat() {
	fmt.Println("dog eat...")
}
func (c cat) eat() {
	fmt.Println("cat eat...")
}

func main() {
	nokia := mobile{}
	nokia.playVideo()
	nokia.playMusic()

	var pet peter
	pet = dog{}
	pet.eat()
	pet = cat{}
	pet.eat()

	var p player
	p = mobile{}
	var v video
	v = mobile{}
	p.playMusic()
	v.playVideo()
}
