package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"../config"
	"../matrix"
	"../utils"
	"gonum.org/v1/gonum/mat"
)

var help bool
var dim int
var iter int
var datasize int

/********************task paramters*************************/

func init() {
	flag.BoolVar(&help, "h", false, "Print help message")
	flag.IntVar(&dim, "d", 10, "The dimension of the matrix")
	flag.IntVar(&iter, "i", 100, "The iteration number of process")
	flag.IntVar(&datasize, "s", 1000, "data size in KB")
}

//client sends task request to scheduler by coap or http
//client gets the assigned task server info
//client sends the task info to task server and time the finish duration
func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	var err error

	var Debug = config.Debug

	//task server parameters
	serverAddr := []byte{}
	serverHTTPPort := config.ServerHTTPPort
	serverPath := config.ServerPath

	//task scheduler parameters
	schHTTPPort := config.SchHTTPPort
	schAddr := config.SchAddr
	schPath := config.SchPath

	/***********************task genertaion***************************/
	//Generate the task contents
	// Initialize two matrices, a and b.
	m1 := matrix.RandomMatrix(dim)
	m2 := matrix.RandomMatrix(dim)

	enc1, err := m1.MarshalBinary()
	if err != nil {
		utils.CheckErr(err, "HTTP POST error")
	}
	enc2, err := m2.MarshalBinary()
	if err != nil {
		utils.CheckErr(err, "HTTP POST error")
	}
	taskDataJSON, err := json.Marshal(map[string][]byte{"iter": utils.Uint32ToBytes(uint32(iter)), "m1": enc1, "m2": enc2})
	if err != nil {
		utils.CheckErr(err, "HTTP POST error")
	}
	/***********************task genertaion***************************/

	/***********************send task parameters to scheduler***********************/
	rand.Seed(time.Now().UnixNano())
	schData := map[string]int{"id": rand.Intn(1000000), "tasksize": dim, "taskiter": iter, "datasize": datasize}
	schDataJSON, _ := json.Marshal(schData)
	reqSch := bytes.NewBuffer(schDataJSON)
	var respSch *http.Response
	respSch, err = http.Post("http://"+string(schAddr)+":"+string(schHTTPPort)+schPath, "application/json;charset=utf-8", reqSch)
	if err != nil {
		utils.CheckErr(err, "Scheduler HTTP POST error")
	}

	bodySch, err := ioutil.ReadAll(respSch.Body)
	if err != nil {
		utils.CheckErr(err, "Scheduler HTTP response error")
	}
	respSch.Body.Close()
	var resultSch map[string][]byte
	err = json.Unmarshal(bodySch, &resultSch)
	if err != nil {
		utils.CheckErr(err, "Scheduler response json decode error")
	} else {
		serverAddr = resultSch["serverAddr"]
	}
	/***********************send task parameters to scheduler***********************/

	/***********************Send to task server and timing***************************/
	//Start to time
	start := time.Now()
	req := bytes.NewBuffer(taskDataJSON)
	var resp *http.Response
	resp, err = http.Post("http://"+string(serverAddr)+":"+string(serverHTTPPort)+serverPath, "application/json;charset=utf-8", req)
	if err != nil {
		utils.CheckErr(err, "HTTP POST error")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.CheckErr(err, "HTTP POST error")
	}
	resp.Body.Close()

	var result map[string][]byte
	err = json.Unmarshal(body, &result)
	if err != nil {
		utils.CheckErr(err, "HTTP POST error")
	} else {
		var matrixRes mat.Dense
		matrixRes.UnmarshalBinary(result["result"])
		if err != nil {
			utils.CheckErr(err, "HTTP POST error")
		} else {
			// Print the result using the formatter.
			if Debug == true {
				fc := mat.Formatted(&matrixRes, mat.Prefix("    "), mat.Squeeze())
				fmt.Printf("c = %v\n", fc)
			}

		}
	}
	//end of timing, microseconds
	cost := time.Since(start).Nanoseconds() / 1000
	fmt.Println(cost)
	/***********************Send to task server and timing***************************/

	//将结果以CVS格式写入文件
	f, err := os.OpenFile("delay.csv", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		utils.CheckErr(err, "delay.csv file create failed.")
	}
	content := []byte(strconv.Itoa(dim) + "," + strconv.Itoa(iter) + "," + strconv.Itoa(1000) + "," + strconv.FormatInt(cost, 10) + "\n")
	_, err = f.Write(content)
	if err != nil {
		utils.CheckErr(err, "File write error.")
	} else {
		fmt.Println("write file successful")
	}

}
