package main

/* Q3: 這是 Q1 的延續（請將完成的code貼過來使用）
 * 請加入 select, 所以如果 client 段 10 秒內沒回應就將他切斷
 */

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var (
	i int
)

func main() {
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

		go handleConn(conn)

		i++
		println("client connected, count: ", i)
	}
}

func handleConn(c net.Conn) {
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
