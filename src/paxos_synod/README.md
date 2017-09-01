## Paxos synod
-------------------------
这是根据[Lamport-paxos](http://research.microsoft.com/users/lamport/pubs/lamport-paxos.pdf)论文第2章的The Single Decree Synod的基本描述略微修改的Go语言实现版本。
主要用到的技术分别是 [MonogoDB](https://docs.mongodb.com/manual/), [Protobuf](https://developers.google.com/protocol-buffers/), [gRPC](https://grpc.io/docs/quickstart/go.html)，具体的使用可以参考各自的官网。

下载后使用需要先安装MongoDB，并下载相关依赖
<code>go get gopkg.in/mgo.v2</code>
<code>go get gopkg.in/mgo.v2/bson</code>
<code>go get google.golang.org/grpc</code>
<code>go get code.google.com/p/goprotobuf/{proto,protoc-gen-go}</code>
<code>go install  code.google.com/p/goprotobuf/proto</code>

####Paxos具体过程
------------------------
只有1个角色，Priest，一共3个节点，也就是一共3个Priest。 Priest可以进行提交提案，决定提案是否通过。提案中包括提案编号和法令。
设定Priest在Chamber中办公，由POST请求至3个Priest中某个Chamber作为接收一个新的意见，Priest可以决定将该意见作为一个新的提案与其他Priest共识之后作为一个正式的法令。POST请求格式要求为 {"decree": "the content of decree"}。

To be continue.