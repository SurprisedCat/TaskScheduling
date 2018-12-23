package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../badgerDB"
	"../config"
	"../converter"
	"../utils"

	"github.com/gorilla/mux"
)

//Debug false default
var Debug = config.Debug
var schHTTPPort = config.SchHTTPPort
var chanSch = make(chan int)
var clientInfo []map[string]int
var assigning map[int][]byte

func main() {

	go scheduling()

	// register router
	router := mux.NewRouter()
	router.
		HandleFunc("/v1/taskschduler/matrixscheduling/post", ClientTaskReq).
		Methods("POST")
	router.HandleFunc("/v1/taskschduler/servreport/post", ServerInfo).
		Methods("POST")

	// start server listening
	err := http.ListenAndServe(":"+string(schHTTPPort), router)
	if err != nil {
		log.Fatalln("ListenAndServe err:", err)
	}

	log.Println("Server end")
}

//ClientTaskReq get task info from client and response the server info
func ClientTaskReq(w http.ResponseWriter, r *http.Request) {
	// parse JSON body
	var req map[string]int
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &req)
	if Debug == true {
		fmt.Println(req)
	}

	/*********************store client info *****************/
	tempinfo := map[string]int{"id": req["id"], "tasksize": req["tasksize"], "taskiter": req["taskiter"], "datasize": req["datasize"]}
	clientInfo = append(clientInfo, tempinfo)
	/**********get the server info***********/
	<-chanSch
	server := assigning[req["id"]]

	// composite response body
	var response = map[string][]byte{"status": []byte("succ"), "server": server}
	responseJSON, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

//ServerInfo get the server info of ip+port CPUHz cpus datarate
func ServerInfo(w http.ResponseWriter, r *http.Request) {
	var req []string
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &req)
	if Debug == true {
		fmt.Println(req)
	}
	if err != nil {
		utils.CheckErr(err, "ServerInfo JSON decode err")
		return
	}
	/****************store the server data into a database redis or leveldb**********************/
	info, err := converter.StringSlice2Cvs(req[1:])
	if err != nil {
		utils.CheckErr(err, "ServerInfo CSV encode err")
		return
	}
	serverID := []byte(req[0])
	err = badgerDB.Set(serverID, []byte(info))
	if err != nil {
		utils.CheckErr(err, "ServerInfo DB write error")
		return
	}
	// composite response body
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Succeed"))

	serverRes := badgerDB.GetAll()
	for _, v0 := range serverRes {
		for k, v := range v0 {
			fmt.Printf("k=%s,v=%s\n", k, v)

		}

	}
	return
}

func scheduling() {

	assigning[clientInfo[0]] = []byte("127.0.0.1:8080")
	chanSch <- 1
}
