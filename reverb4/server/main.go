package main

/* Q4: 這是 Q3 的延續（請將完成的code貼過來使用）
 * 請加入一個取消機制, 如果在主機輸入 .exit 就將程式結束
 */

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var (
	i int
)

func main() {

	abort := make(chan struct{})
	go func() {
		for {
			var s string
			fmt.Scanf("%s", &s)
			if s == ".exit" {
				abort <- struct{}{}
			}

		}
	}()

	i = 0
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			//
		}

		go handleConn(conn, abort)

		i++
		println("client connected, count: ", i)

	}
}

func handleConn(c net.Conn, abort chan struct{}) {
	defer c.Close()

	lines := make(chan string)
	go scan(c, lines)

	timeout := 10 * time.Second
	timer := time.NewTimer(10 * time.Second)
	for {
		select {
		case <-lines:
			timer.Reset(timeout)
		case <-timer.C:
			fmt.Println("timeout。。。。")
			i--
			println("client disconnected, count: ", i)
			return
		case <-abort:
			os.Exit(1)
			return
		default:
			_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
			if err != nil {
				i--
				println("client disconnected, count: ", i)
				return
			}
			time.Sleep(1 * time.Second) // 每秒連續報時！
		}

	}
}

func scan(c net.Conn, lines chan string) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		lines <- input.Text()
	}

}
