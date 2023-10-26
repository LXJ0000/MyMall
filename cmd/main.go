package main

import (
	"MyMall/config"
	"fmt"
)

func main() {
	config.Init()
	fmt.Println(config.AppMode)
}
