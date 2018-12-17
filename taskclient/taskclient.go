package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//client sends task request to scheduler by coap or http
//client gets the assigned task server info
//client sends the task info to task server and time the finish duration
func main() {
	var err error
	var resp *http.Response
	req := bytes.NewBuffer(dataJSON)

	if bytes.Compare(httpPort, []byte("18080")) == 0 {
		resp, err = http.Post("http://"+string(serverAddr)+":"+string(httpPort)+"/v1/upload/aggre", "application/json;charset=utf-8", req)

	} else {
		resp, err = http.Post("http://"+string(serverAddr)+":"+string(httpPort)+"/v1/upload/single", "application/json;charset=utf-8", req)

	}
	if err != nil {
		utils.CheckErr(err, "HTTP POST error")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.CheckErr(err, "HTTP POST error")
	}
	resp.Body.Close()
	httpResp := map[string][]byte{}
	json.Unmarshal(body, &httpResp)
	if bytes.Compare(httpResp["status"], []byte("OK")) == 0 {
		fmt.Printf("http response : %s\n", httpResp)
	} else {
		fmt.Printf("http response : %s\n", httpResp)
	}
}
