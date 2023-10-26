package main

import (
	"MyMall/config"
	"MyMall/route"
)

func main() {
	config.Init()

	r := route.NewRoute()
	_ = r.Run(config.HttpPort)
}
