package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	emptyctx := context.Background()
	cancelctx, cancelfunc := context.WithCancel(emptyctx)

	go DoLongWork(cancelctx)

	var input string
	for {
		fmt.Scanln(&input)
		if input == "done" {
			cancelfunc()
			break
		}
	}
	time.Sleep(200 * time.Millisecond)
	fmt.Println(cancelfunc)
}

func DoLongWork(c context.Context) {
	for {
		fmt.Println("Doing some work ...")
		time.Sleep(300 * time.Millisecond)

		select {
		case <-c.Done():
			fmt.Println("We got canceled by another goroutine")
			goto Finished
		default:
			fmt.Println("Still doing some work ...")

		}
	}
Finished:
}
