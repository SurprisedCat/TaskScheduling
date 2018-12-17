package config

import (
	"encoding/json"
)

//SchAddr Scheduling proxy Address 任务处理服务端地址
var SchAddr = []byte("127.0.0.1")

//HttpPort 服务端的端口号
var HttpPort = []byte("8000")

//PayloadData 具体的payload信息
var PayloadData = []byte("I am OAI cx")

//PayloadJSON 实际负载
var PayloadJSON []byte

//初始化整体payload信息，不要改这里
func init() {

	PayloadJSON, _ = json.Marshal(map[string][]byte{"data": PayloadData})

}
