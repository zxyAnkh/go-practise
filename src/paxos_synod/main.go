package main

import (
	"./config"
	"./core"
	"flag"
	"fmt"
	"strings"
)

var (
	id         int              = 0
	configPath string           = ""
	nodes      []*core.NodeInfo = make([]*core.NodeInfo, 0)
)

func main() {
	parseArguments()
	paxosConfig, err := config.GetPaxosConfig(configPath)
	if err != nil {
		fmt.Errorf("error: %v\n", err)
	}
	nodes = make([]*core.NodeInfo, len(paxosConfig.Paxos.Node))
	for k, v := range paxosConfig.Paxos.Node {
		nstrs := strings.Split(v, ":")
		nodes[k] = core.NewNodeInfo(k+1, nstrs[0], nstrs[1], nstrs[2], nstrs[3])
	}
	chamber := core.NewChamber()
	go chamber.StartServer(nodes[id-1].Ip, nodes[id-1].ServerPort)
	chamber.StartHttpServer(id, nodes[id-1].Ip, nodes[id-1].HttpPort)
}

func parseArguments() {
	flag.StringVar(&configPath, "configPath", "~/config.yaml", "Paxos Config File")
	flag.IntVar(&id, "id", 0, "node id")
	flag.Parse()
}
