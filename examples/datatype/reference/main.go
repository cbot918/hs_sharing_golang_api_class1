package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	/*** array ***/
	// 宣告 array
	array := []int{1, 2, 3, 4, 5}
	fmt.Println("array:")
	fmt.Println(array)

	// 加入元素
	array = append(array, 6)
	fmt.Println(array)

	array[0] = -1
	fmt.Println(array)

	/*** map ***/
	// 宣告 map
	obj := make(map[string]any)
	obj["a"] = 1

	obj2 := make(map[string]any)
	obj2["c"] = 2

	obj["b"] = obj2

	fmt.Println("map:")
	printJson(obj)

}

func printJson(v any) {
	json, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("json marshal error: ", err)
	}

	fmt.Println(string(json))
}
