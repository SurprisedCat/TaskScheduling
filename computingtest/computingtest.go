package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"../matrix"
	"../utils"
	"gonum.org/v1/gonum/mat"
)

var help bool
var dim int
var iter int
var cpuHz string
var cpus string

func init() {

	flag.BoolVar(&help, "h", false, "Print help message")
	flag.IntVar(&dim, "d", 100, "The dimension of the test matrix")
	flag.IntVar(&iter, "i", 100, "The iteration number of the test")
	flag.StringVar(&cpuHz, "z", "1000", "The frequency of CPU measured in MHz")
	flag.StringVar(&cpus, "c", "0.5", "The cpu share of the container")

}
func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	iter := iter
	dim := dim
	// Initialize two matrices, a and b.

	a := matrix.RandomMatrix(dim)
	b := matrix.RandomMatrix(dim)

	// Take the matrix product of a and b and place the result in c.
	var c mat.Dense
	start := time.Now()
	for i := 0; i < iter; i++ {
		c.Mul(a, b)
	}
	cost := time.Since(start).Nanoseconds() / 1000 //MicroSeconds

	//将结果以CVS格式写入文件
	f, err := os.OpenFile("delayTest.csv", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		utils.CheckErr(err, "delay.csv file create failed.")
	}
	content := []byte(cpus + "," + cpuHz + "," + strconv.Itoa(dim) + "," + strconv.Itoa(iter) + "," + strconv.FormatInt(cost, 10) + "\n")
	_, err = f.Write(content)
	if err != nil {
		utils.CheckErr(err, "File write error.")
	} else {
		fmt.Println("write file successful")
	}
}
