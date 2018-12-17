package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"../config"
	"../utils"
)

//client sends task request to scheduler by coap or http
//client gets the assigned task server info
//client sends the task info to task server and time the finish duration
func main() {

	httpPort := config.HttpPort
	dataJSON := config.PayloadJSON
	serverAddr := config.SchAddr
	var err error
	var resp *http.Response
	req := bytes.NewBuffer(dataJSON)

	resp, err = http.Post("http://"+string(serverAddr)+":"+string(httpPort)+"/", "application/json;charset=utf-8", req)

	if err != nil {
		utils.CheckErr(err, "HTTP POST error")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.CheckErr(err, "HTTP POST error")
	}
	resp.Body.Close()
	fmt.Printf("%s\n", body)

	//GET TEXT
	resp, err = http.Get("http://" + string(serverAddr) + ":" + string(httpPort) + "/text")
	if err != nil {
		utils.CheckErr(err, "HTTP POST error")
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.CheckErr(err, "HTTP POST error")
	}
	resp.Body.Close()
	fmt.Printf("%s\n", body)
}
