package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var day int = now.Day()
	var month time.Month = now.Month()
	var year int = now.Year()

	fmt.Println(day)
	fmt.Println(month)
	fmt.Println(year)
}
