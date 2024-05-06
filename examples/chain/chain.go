package main

import "fmt"

type From struct {
	x int
}

func NewFrom(x int) *From {
	return &From{
		x: x,
	}
}

func (f *From) then(fn func(x int) int) *From {
	return NewFrom((f.x))
}

func (f *From) log() {
	fmt.Println(f.x)
}

func main() {

	addOne := func(x int) int {
		return x + 1
	}

	NewFrom(5).
		then(addOne).
		log()

}
