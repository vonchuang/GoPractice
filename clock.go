/******************************************
$ go build ./clock.go
$ ./clock UTC=8000
$ nc localhost 8000 (multiple clients)
******************************************/

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	host := make(map[string]string)

	for _, hosts := range os.Args[1:] {

		var i int = 0
		var tmp string
		for _, s := range strings.Split(hosts, "=") {
			if i == 0 {
				tmp = s
				i++
			} else {
				host[tmp] = s
			}
		}
	}

	for name, id := range host {
		fmt.Printf("%s: %s\n", name, id)
	}

	for name, id := range host {
		listener, err := net.Listen("tcp", "localhost:"+id)
		// fmt.Printf("listen:%T\n", listener) // *net.TCPListener
		if err != nil {
			log.Fatal(err)
		}

		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Print(err)
				continue
			}
			go handleConn(conn, name)
		}

	}

}

func handleConn(c net.Conn, local string) {

	fmt.Printf("%s\n", local)
	defer c.Close()
	loc, err := time.LoadLocation(local)
	if err != nil {
		fmt.Println(err)
	}
	for {
		_, err := io.WriteString(c, time.Now().In(loc).Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
