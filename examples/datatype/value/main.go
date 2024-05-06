package main

import "fmt"

func main() {

	// int
	var number1 int
	number2 := 3

	// string
	var string1 string
	string2 := "hello"

	// bool
	var boolean1 bool
	boolean2 := true

	// struct
	struct1 := struct {
		name  string
		email string
	}{
		name:  "yale",
		email: "yale918@gmail.com",
	}

	fmt.Println(number1)
	fmt.Println(number2)
	fmt.Println(string1)
	fmt.Println(string2)
	fmt.Println(boolean1)
	fmt.Println(boolean2)

	fmt.Println(struct1)
}
