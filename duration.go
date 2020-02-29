package main

import (
	"fmt"
	"time"
)

func main() {
	addHoursTime, _ := time.ParseDuration(fmt.Sprintf("%dm", 240/60))
	fmt.Printf("%v\n", addHoursTime)
	fmt.Printf("%v\n", time.Duration(240)*time.Second)
	fmt.Printf("%v\n", time.Duration(240)*time.Second == addHoursTime)
}
