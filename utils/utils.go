package utils

import (
	"encoding/binary"
	"fmt"
	"os"
)

/*
ClinetData parameters for main function
*/
type ClinetData struct {
	id         string
	dataOrigin string
}

//ClientParameter = Parameter{"127.0.0.1","I am a phd"}

/*
CheckErr 错误处理函数
*/
func CheckErr(err error, extra string) bool {
	if err != nil {
		formatStr := " Err : %s\n"
		if extra != "" {
			formatStr = extra + formatStr
		}

		fmt.Fprintf(os.Stderr, formatStr, err.Error())
		return true
	}

	return false
}

/*
GetClientID get the hostname and uid of the user
*/
func GetClientID(v ...string) []byte {
	//Get the hostname of the machine
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	item := ""
	for _, value := range v {
		item += value
	}
	//clientID = hostname + uid
	clientID := hostname + item
	return []byte(clientID)
}

//Min get the smaller one
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//Uint32ToBytes
func Uint32ToBytes(i uint32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func BytesToUint32(buf []byte) uint32 {
	return uint32(binary.BigEndian.Uint32(buf))
}
