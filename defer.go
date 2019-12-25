package main

import(
	"fmt"
)

func main(){
	func f(x int)(result int){
		defer func (){
			fmt.Printf("Into defer,x = %d\n", x)
			result += x
		}
		return x+x
	}

	fmt.Printf("x=%d\n", f(4))

}
