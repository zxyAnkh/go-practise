package main

import (
	"./config"
	"./core"
	"flag"
	"fmt"
	"strings"
)

type cmdArgs struct {
	ConfigPath string
	Id         int
}

var (
	id    int
	nodes []*core.NodeInfo
)

func main() {
	argus := parseArguments()
	id = argus.Id
	paxosConfig, err := config.GetPaxosConfig(argus.ConfigPath)
	if err != nil {
		fmt.Errorf("error: %v\n", err)
	}
	nodes = make([]*core.NodeInfo, len(paxosConfig.Paxos.Node))
	for k, v := range paxosConfig.Paxos.Node {
		nstrs := strings.Split(v, ":")
		nodes[k] = core.NewNodeInfo(k+1, nstrs[0], nstrs[1], nstrs[2])
	}
	chamber := core.NewChamber()
	chamber.StartServer(nodes[id-1].Ip, nodes[id-1].ServerPort)
}

func parseArguments() cmdArgs {
	argus := cmdArgs{}
	var configPath string
	flag.StringVar(&configPath, "configPath", "~/config.yaml", "Paxos Config File")
	var id int
	flag.IntVar(&id, "id", 0, "node id")
	flag.Parse()
	argus.ConfigPath = configPath
	argus.Id = id
	return argus
}
