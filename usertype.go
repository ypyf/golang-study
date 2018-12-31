package main

type MyInt int

func (i MyInt) Add(a int) MyInt {
	return MyInt(int(i) + a)
}

type F1 func()

func (f F1) Call() {
	println("Before calling a function")
	f()
	println("After calling a function")
}

func main() {
	a := MyInt(10)
	b := a.Add(15)
	println(b)

	f := func() {
		println("Call a function")
	}
	F1(f).Call()
}
