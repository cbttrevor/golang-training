package main

import (
	"fmt"
	"time"
)

func main() {
	onehour, _ := time.ParseDuration("3605s")
	fmt.Println(onehour.Seconds())
	thirtymin, _ := time.ParseDuration("30m")
	fmt.Println(onehour.Truncate(thirtymin))
	fivemin, _ := time.ParseDuration("5m")
	fmt.Println(fivemin)
	fivekns, _ := time.ParseDuration("5000ns")
	fmt.Println(fivekns)

	minusfivemin, _ := time.ParseDuration("6m")
	fmt.Println(minusfivemin.Abs())

	time.Sleep(time.Second * 5)
}
