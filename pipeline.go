/***************************************

$ go run pipeline.go

***************************************/

package main

import(
	"fmt"
)

func main(){
	naturals := make(chan int)
	squares := make(chan int)

	go func(){
		for i:= 0; i < 10; i++{
			naturals <- i;
		}
		close(naturals)	//if not: fatal error: all goroutines are asleep - deadlock!
	}()		//if no (): expression in go must be function call 

	go func(){
		for i := range naturals{
			squares <- i*i
		}
		close(squares)
	}()

	for i := range squares{
		fmt.Printf("%d\t", i)
	}

}
