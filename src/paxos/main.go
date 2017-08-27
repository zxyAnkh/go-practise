package main

import (
	"./config"
	"flag"
	"fmt"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "configPath", "~/config.yaml", "Paxos Config File")
	flag.Parse()
	paxosConfig, err := config.GetPaxosConfig(configPath)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}
	fmt.Printf("%v\n", paxosConfig)
}
