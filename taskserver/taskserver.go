package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"../config"
	"../utils"

	"github.com/gorilla/mux"
	"gonum.org/v1/gonum/mat"
)

func main() {

	//cpu number (normalized based on 1GHz) and bandwidth Kbps
	serverInfo := map[string]int{"cpus": 1, "bw": 500}
	infoJSON, _ := json.Marshal(serverInfo)
	//Role as a server
	//Report info abot itself to scheduler
	go ReportInfo(infoJSON)

	// register router
	router := mux.NewRouter()
	router.HandleFunc("/v1/taskserver/{servicename}/post", MatrixComputing).
		Methods("POST")

	// start server listening
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatalln("ListenAndServe err:", err)
	}
	log.Println("Server end")

}

//MatrixComputing process
func MatrixComputing(w http.ResponseWriter, r *http.Request) {
	// parse path variable
	vars := mux.Vars(r)
	servicename := vars["servicename"]

	// parse JSON body
	var req map[string][]byte
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &req)

	//Executing the service
	//Start to time
	start := time.Now()
	var dec1 mat.Dense
	dec1.UnmarshalBinary(req["m1"])
	var dec2 mat.Dense
	dec2.UnmarshalBinary(req["m2"])
	iter := utils.BytesToUint32(req["iter"])
	var res mat.Dense
	var i uint32
	for i = 0; i < iter; i++ {
		res.Mul(&dec1, &dec2)
	}
	//end of timing, microseconds
	cost := time.Since(start).Nanoseconds() / 1000
	fmt.Println(cost)

	resJSON, err := res.MarshalBinary()
	if err != nil {
		utils.CheckErr(err, "HTTP POST error")
	}
	// composite response body
	var response = map[string][]byte{"status": []byte("succ"), "name": []byte(servicename), "result": resJSON}
	responseJSON, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

//ReportInfo report information of cpu and bandwidth to scheduler
func ReportInfo(infoJSON []byte) {
	for {
		//task scheduler parameters
		schHttpPort := config.SchHttpPort
		schAddr := config.SchAddr
		schPath := config.SchPathServerReport

		reqSch := bytes.NewBuffer(infoJSON)
		var respSch *http.Response
		respSch, err := http.Post("http://"+string(schAddr)+":"+string(schHttpPort)+schPath, "application/json;charset=utf-8", reqSch)
		if err != nil {
			utils.CheckErr(err, "Scheduler HTTP POST error")
		}

		bodySch, err := ioutil.ReadAll(respSch.Body)
		if err != nil {
			utils.CheckErr(err, "Scheduler HTTP response error")
		}
		respSch.Body.Close()
		if bytes.Compare(bodySch, []byte("Succeed")) != 0 {
			log.Println("Report info is not received correctly")
			fmt.Println(bodySch)
		} else {
			log.Println("Report OK")
		}
		time.Sleep(time.Duration(time.Second * 2))
	}
}
