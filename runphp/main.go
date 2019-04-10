package main

import (
	"fmt"
	"log"
	"os/exec"
)

func RunPHP(code string) (string, error) {
	cmd := exec.Command("php")
	p, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalln(err)
	}
	_, err = p.Write([]byte(code))
	if err != nil {
		log.Fatalln(err)
	}

	if err = p.Close(); err != nil {
		log.Fatalln(err)
	}
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func main() {
	out, err := RunPHP(`<?php
		echo phpinfo();
		$a = 123;
		$b = 456;
		$c = $a * $b;
		echo "$c\n";
	?>`)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(out)
}
