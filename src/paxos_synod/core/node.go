package core

type NodeInfo struct {
	Id         int
	Ip         string
	ClientPort string
	ServerPort string
	HttpPort   string
}

func NewNodeInfo(id int, ip, clientPort, serverPort, httpPort string) NodeInfo {
	return NodeInfo{
		Id:         id,
		Ip:         ip,
		ClientPort: clientPort,
		ServerPort: serverPort,
		HttpPort:   httpPort,
	}
}
