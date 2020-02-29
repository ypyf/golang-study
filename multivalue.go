package main

func multivalue() (int, int) {
	return 1, 2
}

func main() {
	var a map[string]int
	a = make(map[string]int)
	a["aa"], a["bb"] = multivalue()
	for k := range a {
		println(k)
	}
}
