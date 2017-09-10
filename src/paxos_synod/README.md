## Paxos synod
-------------------------
这是根据[Lamport-paxos](https://www.microsoft.com/en-us/research/wp-content/uploads/2016/12/The-Part-Time-Parliament.pdf)论文第2章的The Single Decree Synod的基本描述略微修改的Go语言实现版本。
主要用到的技术分别是 [MonogoDB](https://docs.mongodb.com/manual/), [Protobuf](https://developers.google.com/protocol-buffers/), [gRPC](https://grpc.io/docs/quickstart/go.html)，具体的使用可以参考各自的官网。

下载后使用需要先安装MongoDB，并下载相关依赖
```
$go get gopkg.in/mgo.v2
$go get gopkg.in/mgo.v2/bson
$go get google.golang.org/grpc
$go get code.google.com/p/goprotobuf/{proto,protoc-gen-go}
$go install code.google.com/p/goprotobuf/proto
```

####Paxos中数据结构的定义
------------------------
######数据库 数据结构
```
type Leger{
    Id uint32 // the ballot id
    Decree string // the decree in this ballot
    Priest int // the priest who begin this ballot
}
type Note{
    Id uint32 // the ballot id
    Decree string // the decree in this ballot
    Priest int // the priest who begin this ballot
}
```
在该版本中，由于都部署在本地，所以使用同一个数据库。
######Messages
```
message NextBallot{
    uint32 id = 1; // the ballot id
    uint32 priest = 2; // the priest who send this request
}

message LastVote{
    uint32 id = 1; // the ballot id
    uint32 maxId = 2; // the max ballot id less than id of this priest
    uint32 priest = 3; // the priest who send this request
}

message BeginBallot{
    uint32 id = 1; // the ballot id
    string decree = 2; // the decree of this ballot
    uint32 priest = 3; // the priest who send this request
}

message Voted{
    bool vote = 1; // the flag whether vote this ballot
    uint32 id = 2; // the ballot id
    uint32 priest = 3; // the priest who send this request
}

message Success{
    uint32 id = 1; // the ballot id
    string decree = 2; // the decree of this ballot
    uint32 priest = 3; // the preist who begin this ballot
}
```

####Paxos具体过程
------------------------
只有1个角色，Priest，一共3个节点，也就是一共3个Priest。 Priest可以进行提交提案，决定提案是否通过。提案中包括提案编号和法令。
设定Priest在Chamber中办公，由POST请求至3个Priest中某个Chamber作为接收一个新的意见，Priest可以决定将该意见作为一个新的Ballot与其他Priest共识之后成为一个正式的Decree记入各自的Leger中。POST请求格式要求为 ```{"decree": "the content of decree"}```。
1. Priest p1收到一个POST请求，p1检查自己的Leger和Note记录，判断是否存在相同内容的记录，存在则忽略此请求，不存在则生成一个新的Ballot id和一个NextBallot请求，并发送给Priest p2和p3。

2. Priest p2收到来自p1的NextBallot消息之后，根据自己Note中的信息，找到自己投票的小于信息NextBallot中Ballot id的最大的Ballot id，并返回LastVote信息给p1，如果没有找到，则返回空的LastVote信息。

3. 当p1收到大部分Priest即p2和p3的回复后，将该Ballot id的Ballot的Decree改为遵守Paxos协议的decree，并生成一个BeginBallot信息，将其发送给其他的Priest。
    遵守Paxos协议即，如果收到p2和p3的回复后，p2或者p3投票给了之前一个尚未通过的Ballot，则该Ballot的decree修改为小于该Ballot id的最大那个Ballot的Decree。
4. 其他的Priest收到BeginBallot消息后，根据之前它给其他Priest返回的LastVote，决定是否投票给该BeginBallot，如果决定投票，则将其记录在Leger中，并发送Voted信息给p1。

5. 如果p1从所有的Priest的大部分Priest中收到Voted回复，则在他的Leger上记录该decree，并发送一个Success信息给每一个Priest。

6. Priest在收到Success消息后，Priest都将在Leger中记录decree。

To be continue...
