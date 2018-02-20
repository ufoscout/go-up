package main

import (
	"fmt"

	"github.com/ufoscout/go-up"
)

func main() {
	up, err := go_up.NewGoUp().
		AddFile("./config.properties", false).
		Build()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	fmt.Println(up.GetString("hello"))
}
