package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

func main() {
	emptyctx := context.Background()
	funcDuration, _ := time.ParseDuration("2s")
	expiryTime := time.Now().Add(funcDuration)

	deadlinectx, cancelfunc := context.WithDeadline(emptyctx, expiryTime)

	fmt.Println(deadlinectx)
	fmt.Println(cancelfunc)
	fmt.Println(reflect.TypeOf(deadlinectx))

	go WorkerCode(deadlinectx)

	time.Sleep(3 * time.Second)
	fmt.Println("Exiting main Goroutine")
}

func WorkerCode(c context.Context) {
	for {
		fmt.Println("Doing some work before deadline ...")
		time.Sleep(300 * time.Millisecond)
		select {
		case <-c.Done():
			fmt.Println("Exceeded deadline, aborting work")
			goto Finished
		default:
			fmt.Println("Doing more work before deadline ...")
		}
	}
Finished:
}
