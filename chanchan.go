package main

import (
	"fmt"
)

type client chan string

func main() {
	ch := make(chan client)
	mssg := make(chan string)
	go func() {
		mssg <- "Hello World"
	}()

	go func() {
		ch <- mssg
	}()

	out := <-ch

	fmt.Println(<-out)
}
