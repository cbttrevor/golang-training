package main

import (
	"fmt"
	"time"
)

func main() {
	timeLou, _ := time.LoadLocation("America/Louisville")
	fmt.Println(timeLou)
	timeLA, _ := time.LoadLocation("America/Los_Angeles")
	fmt.Println(timeLA)

	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Println(currentTime.In(timeLou))
	fmt.Println(currentTime.In(timeLA))

	laNoon := time.Date(2022, time.August, 5, 12, 0, 0, 0, timeLA)
	fmt.Println(laNoon)
	fmt.Println(laNoon.In(timeLou))

}
