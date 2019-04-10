package main

func main() {
	a := make(chan int, 1)
	b := make(chan string, 1)

	select {
	case x := <-a:
		println(x)
	case y := <-b:
		println(y)
	}
}

// go func() {
// 	a <- 1
// }()

// go func() {
// 	b <- "hello"
// }()
