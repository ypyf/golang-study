package main

import "fmt"

type Item interface {
	ShowPrice() float32
	ShowTitle() string
}

type Book struct {
	title    string
	price    float32
	quantity int32
}

type Food struct {
	title    string
	price    float32
	quantity int32
}

type Car struct{}

// *Book实现了Item接口

func (b *Book) ShowPrice() float32 {
	return b.price
}

func (b *Book) ShowTitle() string {
	return b.title
}

// *Food实现了Item接口

func (f *Food) ShowPrice() float32 {
	return f.price
}

func (f *Food) ShowTitle() string {
	return f.title
}

func makeItemDesc(item Item) string {
	return fmt.Sprintf("您选择的商品是 %s, 价格 %.2f",
		item.ShowTitle(), item.ShowPrice())
}

func main() {
	b1 := &Book{"Programming in Go", 33.4, 10}
	f1 := &Food{"西瓜", 50.0, 10}
	fmt.Println(makeItemDesc(b1))
	fmt.Println(makeItemDesc(f1))

	// 字段的类型、名称和数量相同的类型之间可以转换
	b2 := Book(*f1)
	b3 := (*Book)(f1) // 指针亦可转换
	fmt.Println(makeItemDesc(&b2))
	fmt.Println(makeItemDesc(b3))

	// 类型断言,点(.)左边必须是一个接口类型的变量
	b4 := Item(f1).(*Food) // 转换失败将出现异常
	fmt.Println(b4)

	// 检查类型断言
	b5, ok := Item(f1).(*Food)
	if ok {
		fmt.Println(b5)
	} else {
		fmt.Println("类型断言失败")
	}
}
