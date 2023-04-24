package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime.Format("02 01 03:04 2006"))
	duration, _ := time.ParseDuration("-43824h")
	oldTime := currentTime.Add(duration)
	fmt.Println(oldTime.Format("02 01 03:04 2006"))
}
