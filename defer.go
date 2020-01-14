package main

func dosomething(a string, c int) {
	println(a)
}

func inc(i *int) int {
	(*i)++
	return *i
}

func main() {
	a := 0
	defer dosomething("hello", inc(&a))
	println(a)
}
