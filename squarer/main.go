package main

// Q：請改寫main_original.go 的 function，讓程式可以依以下方式繼續執行
// hint：把 chan 做成 output

/************* 請修改以上範圍 **************** */
func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func squarer(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func printer(in <-chan int) {
	for square := range in {
		println(square)
	}
}

/*************	請修改以上範圍 ************ */

func main() {

	c := gen()
	s := squarer(c)

	printer(s)
}
