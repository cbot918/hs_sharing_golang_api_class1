package main

import "fmt"

func main() {

	addTwo := func(x int) int {
		return x + 2
	}

	minusOne := func(x int) int {
		return x - 1
	}

	calculate := func(x int, fn func(int) int) int {
		return fn(x)
	}

	ans := 50

	ans = calculate(ans, addTwo)
	ans = calculate(ans, minusOne)

	fmt.Println(ans)

}
