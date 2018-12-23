package config

import (
	"encoding/json"
)

//Debug for verbose print
var Debug = false

//DataPath path of the database
var DataPath = "./"

/*******************scheduler parameters*********************/

//SchAddr Scheduling proxy Address 任务处理服务端地址
var SchAddr = []byte("127.0.0.1")

//SchHTTPPort 服务端的端口号
var SchHTTPPort = []byte("8118")

//SchPath the http server path of scheduling
var SchPath = "/v1/taskschduler/matrixscheduling/post"

//SchPathServerReport scheduler process the server info reporter
var SchPathServerReport = "/v1/taskschduler/servreport/post"

/*******************server parameters*********************/

//ServerHTTPPort port of task server
var ServerHTTPPort = []byte("8000")

//ServerAddr ip address of task server
var ServerAddr = []byte("127.0.0.1")

//ServerPath Path the path of restful URL used by server
var ServerPath = "/v1/taskserver/matrixcomputing/post"

/*******************client parameters*********************/

//PayloadData 具体的payload信息
var PayloadData = []byte("I am OAI cx")

//PayloadJSON 实际负载
var PayloadJSON []byte

//初始化整体payload信息，不要改这里
func init() {

	PayloadJSON, _ = json.Marshal(map[string][]byte{"data": PayloadData})

}
