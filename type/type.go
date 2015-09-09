package main

import "fmt"

type Item interface {
	ShowPrice()
	ShowTitle()
}

type Book struct {
	title string
	price float32
}

type Food struct {
	title string
	price float32
}

func (b *Book) ShowPrice() {
	fmt.Println(b.price)
}

func (b *Book) ShowTitle() {
	fmt.Println(b.title)
}

func makeDesc(item Item) {

}

func main() {
	b1 := Book{"Programming in Go", 33.4}
	b1.ShowPrice()
	b1.ShowTitle()
}
