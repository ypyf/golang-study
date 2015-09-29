package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var php_code string = `<?php
phpinfo();
`

func RunPHP(code string) error {
	cmd := exec.Command("php")
	p, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	p.Write([]byte(code))
	p.Close()
	out, err := cmd.CombinedOutput()
	if len(out) > 0 {
		fmt.Printf("%s\n", out)
	}
	return err
}

func Repl() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	text, _ := reader.ReadString('\n')
	RunPHP(text)
}
func main() {
	for {
		Repl()
	}
}
