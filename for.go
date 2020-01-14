package main

var a chan int

func endless() <-chan int {
	for {
		a <- 1
	}
}
func main() {
	a = make(chan int)
	go endless()
	for x := range a {
		println(x)
	}
}
