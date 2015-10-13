package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
)

type config struct {
	Redis struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	f, err := os.Open("config.yaml")
	check(err)

	data, err := ioutil.ReadAll(f)
	check(err)

	c := new(config)
	err = yaml.Unmarshal([]byte(data), c)
	check(err)

	fmt.Printf("%+v", c)
}
