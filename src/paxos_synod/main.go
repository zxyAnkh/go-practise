package main

import (
	"./config"
	"./core"
	"flag"
	"fmt"
	"strings"
)

var (
	id         int             = 0
	configPath string          = ""
	nodes      []core.NodeInfo = make([]core.NodeInfo, 0)
)

// start command: go run main.go -id 1 -configPath ~/config.yaml
func main() {
	parseArguments()
	paxosConfig, err := config.GetPaxosConfig(configPath)
	if err != nil {
		fmt.Errorf("error: %v\n", err)
	}
	nodes = make([]core.NodeInfo, len(paxosConfig.Paxos.Node))
	for k, v := range paxosConfig.Paxos.Node {
		nstrs := strings.Split(v, ":")
		nodes[k] = core.NewNodeInfo(k+1, nstrs[0], nstrs[1], nstrs[2], nstrs[3])
	}
	if len(nodes) == 0 {
		fmt.Println("No node info, exit.")
		return
	}
	err = core.InitPriest(id, nodes)
	if err != nil {
		fmt.Printf("Init priest information error: %v\n", err)
		return
	}
	core.InitChamber(nodes[id-1].Ip, nodes[id-1].ServerPort, nodes[id-1].HttpPort)
}

func parseArguments() {
	flag.StringVar(&configPath, "configPath", "~/config.yaml", "Paxos Config File")
	flag.IntVar(&id, "id", 0, "node id")
	flag.Parse()
}
