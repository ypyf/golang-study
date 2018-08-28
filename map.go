package main

func main() {
	a := make(map[string]int)
	a["hello"] = 1
	a["world"] = 2
	for k, v := range a {
		println(k, v)
	}
}
