package main
// 題目請看 main.go

func main() {

	msg := make(chan string)

	strings := []string{"hello", "world", "animal", "planet"}

	for i, s := range strings {
		go func(i int, s string) {
			msg <- s
		}(i, s)
	}

	for i := 0; i < len(strings); i++ {
		println("The message you get is:", <-msg, " original message is: ", strings[i])
	}

	println("conclusion: it does't arrive in the same sequence as sent")

}
