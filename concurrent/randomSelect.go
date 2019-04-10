package main

func Select() int {
	ch := make(chan int, 1)
	select {
	case ch <- 0:
	case ch <- 1:
	case ch <- 2:
	case ch <- 3:
	case ch <- 4:
	case ch <- 5:
	case ch <- 6:
	case ch <- 7:
	case ch <- 8:
	case ch <- 9:
	}
	return <-ch
}

func main() {
	for i := 0; i < 100; i++ {
		print(Select())
	}
}
