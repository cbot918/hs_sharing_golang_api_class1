package util

import (
	"encoding/json"
	"fmt"
)

func PrintJson(v any) {
	json, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("json marshal error: ", err)
	}

	fmt.Println(string(json))
}
