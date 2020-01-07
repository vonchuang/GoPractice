package main

import (
	"fmt"
	"net"
	"strconv"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/netutil"
	"github.com/ethereum/go-ethereum/rpc"
)

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	OK  int    `json:"ok"`
	Msg string `json:"msg"`
}

// CalculatorService ...
type CalculatorService struct{}

// Add ...
func (s *CalculatorService) Add(a, b int) Response {
	fmt.Println("Call Add Function")

	var res Response
	res.OK = 1
	res.Msg = strconv.Itoa(a + b)
	return res
}

// Hello ...
func (s *CalculatorService) Hello(item Item) Response {
	fmt.Println("Call Hello Function")

	var res Response
	res.OK = 1
	res.Msg = "Hello " + item.Name + " " + strconv.Itoa((item.Id))
	return res
}

func main() {
	service := new(CalculatorService)
	// NewServer creates a new server instance with no registered handlers.
	server := rpc.NewServer()
	defer server.Stop()

	// RegisterName creates a service for the given receiver type under the given name.
	if err := server.RegisterName("CalculatorService", service); err != nil {
		fmt.Println("!!CalculatorService!! %s", err)
	}

	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("can't listen:%v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		if netutil.IsTemporaryError(err) {
			log.Warn("RPC accept error", "err", err)
			continue
		}

		fmt.Println("Connect~")

		go func() {
			defer conn.Close()
			log.Trace("Accepted RPC connection", "conn", conn.RemoteAddr())

			codec := rpc.NewCodec(conn)
			go server.ServeCodec(codec, 0)
		}()
	}
}
