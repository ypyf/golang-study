package main

import (
	"log"
	"time"
)

func calculateAge(birthday, thisDay string) int {
	if len(birthday) < 10 || birthday == "0000-00-00" {
		return 0
	}

	date := string([]rune(birthday)[:10])
	t1, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Fatalln(err)
	}

	t2, err := time.Parse("2006-01-02", thisDay)
	if err != nil {
		log.Fatalln(err)
	}

	return t2.Year() - t1.Year()
}

func main() {
	println(calculateAge("1985-11-10 00:00:00", "2018-09-01"))
}
