package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

const (
	BASE64STR = `NDQ5MnlkMzYM0OHhuYzI2UOWgwdWE2ZDWE4cmIyNXc4M2c=\n`
)

func main() {
	// 有的base64编码后的字符串会含有换行符，用下面的方法解码会失败
	data, err := base64.StdEncoding.DecodeString(BASE64STR)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%q\n", data)

	// 只有用下面的方法才能解码成功
	str := make([]byte, 100)
	decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(BASE64STR))
	decoder.Read(str)
	fmt.Printf("%+v\n", str)
}
