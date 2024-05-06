package main

import (
	"app/util"
)

func main() {

	obj := make(map[string]any)

	obj["a"] = 1

	obj2 := make(map[string]any)
	obj2["c"] = 2

	obj["b"] = obj2

	util.PrintJson(obj)

}
