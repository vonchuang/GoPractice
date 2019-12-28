package original

func gen(out chan<- int) {
	for i := 0; i < 20; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for naturals := range in {
		out <- naturals * naturals
	}
	close(out)
}

func printer(in <-chan int) {
	for square := range in {
		println(square)
	}
}

func main() {

	naturals := make(chan int)
	squares := make(chan int)

	go gen(naturals)
	go squarer(squares, naturals)

	printer(squares)

}
