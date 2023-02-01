package main

import "fmt"

type USB interface {
	read()
	write()
}

type Computer struct {
	name string
}

func (c Computer) read() {
	fmt.Printf("%v computer read\n", c.name)
}

func (c Computer) write() {
	fmt.Printf("%v computer ssd\n", c.name)
}

type Mobile struct {
	name string
}

func (m Mobile) read() {
	fmt.Printf("%v mobile read\n", m.name)
}

func (m Mobile) write() {
	fmt.Printf("%v mobile 5g\n", m.name)
}

type Pet interface {
	eat(string) string
}

type Dog struct {
	name string
}

func (dog *Dog) eat(food string) string {
	(*dog).name = "huahua"
	fmt.Printf("%v eat %v\n", (*dog).name, food)
	return "good"
}

type Player interface {
	playMusic()
}

type Video interface {
	playVideo()
}

func (m Mobile) playMusic() {
	fmt.Printf("play music\n")
}

func (m Mobile) playVideo() {
	fmt.Printf("play video\n")
}

type Bird interface {
	fly()
}

type Fish interface {
	swim()
}

type FlyFish interface {
	Bird
	Fish
}

func main() {

	var c0 = Computer{
		name: "AMD",
	}
	c0.read()
	c0.write()

	var m0 = Mobile{"iPhone"}
	m0.read()
	m0.write()
	m0.playMusic()
	m0.playVideo()

	var d0 = Dog{
		name: "wangcai",
	}
	(&d0).eat("pork rib")
}
