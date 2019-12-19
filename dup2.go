// execute: 
// $ go run dup2.go duptext.txt duptest2.txt

package main

import(
	"fmt"
	"os"
	"strings"
	"io/ioutil"
)

func main(){
	
	for _, filename := range os.Args[1:] {
		count := make(map[string]int)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}

		fmt.Printf("%s:\n", filename)

		for _, word := range strings.Split(string(data), "\n"){
			count[word]++
		}

		for word, n := range count{
			if n > 1 {
				fmt.Printf("%d %s\n", n, word);
			}
		}


	}
}
