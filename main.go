package main

import (
	"encoding/json"
	"fmt"

	"github.com/dhavlev/go-object-oriented/business"
)

func main() {
	json, err := json.Marshal(business.BuildOrder())
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json))
}
