package main

import(
	"bufio"
	"fmt"
	"os"
)

func main(){
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		count[input.Text()]++
		
		n := count[input.Text()]
		if n > 1{
			fmt.Printf("%d\t%s\n", n,input.Text())
		}
	}
}
