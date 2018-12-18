package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../config"

	"github.com/gorilla/mux"
)

var Debug = config.Debug
var schHttpPort = config.SchHttpPort
var chanSch = make(chan int)

func main() {

	// register router
	router := mux.NewRouter()
	router.
		HandleFunc("/v1/taskschduler/matrixscheduling/post", ClientTaskReq).
		Methods("POST")
	router.HandleFunc("/v1/taskschduler/servreport/post", ServerInfo).
		Methods("POST")

	// start server listening
	err := http.ListenAndServe(":"+string(schHttpPort), router)
	if err != nil {
		log.Fatalln("ListenAndServe err:", err)
	}

	log.Println("Server end")
}

//ClientTaskInfo get task info from client and response the server info
func ClientTaskReq(w http.ResponseWriter, r *http.Request) {
	// parse JSON body
	var req map[string][]byte
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &req)
	if Debug == true {
		fmt.Println(req)
	}
	/**********executing the scheduling algorithm***********/
	var serverAddr = []byte("127.0.0.1")

	// composite response body
	var response = map[string][]byte{"status": []byte("succ"), "serverAddr": serverAddr}
	responseJSON, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

//ServerInfo
func ServerInfo(w http.ResponseWriter, r *http.Request) {
	var req map[string]int
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &req)
	if Debug == true {
		fmt.Println(req)
	}
	/****************store the server data into a database redis or leveldb**********************/

	// composite response body
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Succeed"))
}
