package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type PaxosConfig struct {
	Paxos Paxos `yaml: "paxos"`
}

type Paxos struct {
	Version string   `yaml: "version"`
	Node    []string `yaml: "node, flow"`
}

func GetPaxosConfig(configPath string) (*PaxosConfig, error) {
	if !fileExists(configPath) {
		return &PaxosConfig{}, fmt.Errorf("Config file not exists, file path is %s\n", configPath)
	}
	configData, err := getFileContent(configPath)
	if err != nil {
		return &PaxosConfig{}, fmt.Errorf("Config file read error: %v\n", err)
	}
	config := &PaxosConfig{}
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return &PaxosConfig{}, fmt.Errorf("Unmarshal config data error: %v\n", err)
	}
	return config, nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func getFileContent(path string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
