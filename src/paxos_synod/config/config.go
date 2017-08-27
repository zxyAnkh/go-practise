package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type PaxosConfig struct {
	Paxos paxos `yaml: "paxos"`
}

type paxos struct {
	Version string   `yaml: "version"`
	Node    []string `yaml: "node, flow"`
}

func GetPaxosConfig(configPath string) (PaxosConfig, error) {
	if !fileExists(configPath) {
		return PaxosConfig{}, fmt.Errorf("Config file not exists, file path is %s", configPath)
	}
	configData, err := getFileContent(configPath)
	if err != nil {
		return PaxosConfig{}, fmt.Errorf("Config file read error: %v", err)
	}
	config := PaxosConfig{}
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return PaxosConfig{}, fmt.Errorf("Unmarshal config data error: %v", err)
	}
	return config, nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func getFileContent(path string) ([]byte, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return []byte(""), fmt.Errorf("Open config file error: %v", err)
	}
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
	return buf, nil
}
