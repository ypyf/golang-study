package main

import (
	"log"
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
	log.Printf("%s\n", out)
	return err
}

func main() {
	RunPHP(php_code)
}
