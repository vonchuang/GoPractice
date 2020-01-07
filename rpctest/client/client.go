package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
)

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	OK  int    `json:"ok"`
	Msg string `json:"msg"`
}

func main() {
	client, err := rpc.DialHTTP("http://127.0.0.1:8888")
	defer client.Close()

	if err != nil {
		log.Fatal("Connect to server fail: ", err)
	}

	var res Response
	if err := client.Call(&res, "testService2_Echo2", "x", 3); err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	/*
		item := Item{1234, "Ni"}
		var res2 Response
		if err := client.Call(&res2, "CalculatorService_Hello", item); err!= nil{
			fmt.Println(err)
		}
		fmt.Println(res2)
	*/
}
