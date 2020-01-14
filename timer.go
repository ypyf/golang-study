package main

import (
	"time"
)

func main() {
	timer := time.NewTimer(0 * time.Second)
	<-timer.C
	timer.Stop()
}
