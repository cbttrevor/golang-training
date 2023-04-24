package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Printf("Is daylight time active? %v\n", currentTime.IsDST())
	fmt.Printf("How many days into the year are we? %v\n", currentTime.YearDay())
	fmt.Printf("How many minutes into the current hour are we? %v\n", currentTime.Minute())
	fmt.Println(currentTime.UTC())
}
