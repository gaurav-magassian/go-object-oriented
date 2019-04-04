package main

import (
	"encoding/json"
	"fmt"

	"./business"
)

func main() {
	json, err := json.Marshal(business.BuildOrder())
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json))
}
