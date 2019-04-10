package main

import "fmt"
import "strings"

func main() {
	// 每个字符类型保存一个Unicode码点
	c := 's'
	fmt.Printf("字符的类型为 %T\n", c)

	// 字符串编码为UTF-8字节
	s := "Microsoft公司"
	fmt.Printf("字符串的类型为 %T\n", s)

	// 无法直接索引非ASCII字符
	fmt.Printf("%c\n", s[0])
	fmt.Printf("%c\n", s[9])

	// 查找字符的位置
	p0 := strings.Index(s, "公")
	p1 := strings.Index(s, "司")
	fmt.Printf("公的位置是 %d\n", p0)
	fmt.Printf("司的位置是 %d\n", p1)
	// 通过[]byte访问string中的汉字
	fmt.Printf("%s\n", s[p0:p1])

	// 迭代字符串
	for i, c := range s {
		fmt.Printf("<%d:%c>", i, c)
	}
	fmt.Println()

	// string转换为[]rune，将utf8编码的字符串转为unicode码点
	// 可以直接索引非ASCII字符
	r := []rune(s)
	fmt.Printf("%c\n", r[10])
}
