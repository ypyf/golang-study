package main

// deadlock
func main() {
	a := make(<-chan int)
	go func() {
		//a <- 1
		//close(a)
	}()

	for i := range a {
		println(i)
	}
}
