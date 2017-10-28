package main

import (
	"./core"
)

func main() {
	server := core.NewServer(8081, 8082)
	server.Start()
}
