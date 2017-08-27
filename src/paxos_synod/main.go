package main

import (
	"./config"
	"flag"
	"fmt"
)

type Arguments struct {
	ConfigPath string
	Id         int
}

func main() {
	arguments := parseArguments()
	paxosConfig, err := config.GetPaxosConfig(arguments.ConfigPath)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}
	fmt.Printf("%v\n", paxosConfig)
}

func parseArguments() Arguments {
	arguments := Arguments{}
	var configPath string
	flag.StringVar(&configPath, "configPath", "~/config.yaml", "Paxos Config File")
	var id int
	flag.IntVar(&id, "id", 0, "node id")
	flag.Parse()
	arguments.ConfigPath = configPath
	arguments.Id = id
	return arguments
}
