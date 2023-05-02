package main

import (
	"context"
	"fmt"
	"reflect"

	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	fmt.Println(ctx)
	fmt.Println(reflect.TypeOf(ctx))

	valuectx := context.WithValue(ctx, "trace-id", uuid.NewString())
	fmt.Println(reflect.TypeOf(valuectx))
	fmt.Println(valuectx)

	DoSomeWork(valuectx)
}

func DoSomeWork(c context.Context) {
	fmt.Printf("Working on the request ID %v\n", c.Value("trace-id"))
}
