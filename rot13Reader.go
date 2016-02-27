package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

const (
	t0 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	t1 = "abcdefghijklmnopqrstuvwxyz"
)

type rot13Reader struct {
	io.Reader
}

// 实现 Read([]byte) (int, error) 接口
func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.Reader.Read(b)
	if err == nil {
		for i := 0; i < n; i++ {
			if unicode.IsUpper(rune(b[i])) {
				j := int8(b[i] - 'A')
				j -= 13
				if j < 0 {
					j = 26 + j
				}
				b[i] = t0[j]
			} else if unicode.IsLower(rune(b[i])) {
				j := int8(b[i] - 'a')
				j -= 13
				if j < 0 {
					j = 26 + j
				}
				b[i] = t1[j]
			}
		}
	}
	return n, err
}

func debuger() {
	var a int16 = 's'
	var b int32 = 'b'
	b = int32(a)
	fmt.Println(a)
	fmt.Println(b)
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
