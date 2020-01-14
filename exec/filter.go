package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// boolInteractive := flag.Bool("c", false, "键盘输入时显示输入的字符.")
	// boolLinefeed := flag.Bool("r", false, "输出的换行是\\r\\n.")
	stringTrailer := flag.String("e", "", "程序结束时输出字符串.")
	flag.Parse()

	name := os.Args[flag.NArg()]
	fmt.Printf("Name %s\n", name)
	fmt.Printf("Rest %v\n", os.Args[flag.NArg()+1:])

	cmd := exec.Command(name, os.Args[flag.NArg()+1:]...)

	bytes, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", []byte{0x1b, 0x1c, 0x08})
		if exitErr, ok := err.(*exec.ExitError); ok {
			fmt.Printf("%s", string(exitErr.Stderr))
		} else {
			fmt.Printf("%s", err)
		}
	} else {
		fmt.Printf("%s", []byte{0x1b, 0x1c, 0x07})
		fmt.Printf("%s", bytes)
	}

	if *stringTrailer != "" {
		fmt.Printf("%s", []byte{0x1b, 0x1c, 0x1f, 0x1e})
		fmt.Printf("%s", *stringTrailer)
	}
}
