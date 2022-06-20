package main

import (
	"encoding/json"
	"fmt"

	"github.com/SharkEzz/provisiond/pkg/loader"
)

func main() {
	cfg, err := loader.GetLoader("./test.yaml").Load()
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(cfg, "", "  ")
	fmt.Println(string(data))
}
