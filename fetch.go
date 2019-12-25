/*************************************

$ go build fetch.go
$ ./fetch http://gopl.io

*************************************/

package main

import(
	"fmt"
	"os"
	"io"
//	"io/ioutil"
	"net/http"
)

func main(){
	for _, url := range os.Args[1:]{
		res, err := http.Get(url)
		if err != nil{
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		if _, err := io.Copy(os.Stdout, res.Body); err != nil{
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		
		/*
		body, err := ioutil.ReadAll(res.Body)
		
		res.Body.Close()
		if err != nil{
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", body)
		*/
	}
}
