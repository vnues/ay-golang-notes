package main

import (
	"context"
	"fmt"
	"time"
)

type paramKey struct{}

func main() {
	c := context.WithValue(context.Background(), paramKey{}, "abc")
	c, cancel := context.WithTimeout(c, 5*time.Second)

	defer cancel()

	mainTask(c)
}

//context在作为参数的时候，按照约定，必须是第一个参数
func mainTask(c context.Context) {
	fmt.Printf("main task started with param %q\n", c.Value(paramKey{}))

	c1, cancel := context.WithTimeout(c, 2*time.Second)
	defer cancel()

	smallTask(c1, "task1", 4*time.Second)
	smallTask(c, "task2", 2*time.Second)
	smallTask(context.Background(), "task3", 3*time.Second)
}

func smallTask(c context.Context, name string, d time.Duration) {
	fmt.Printf("%s started with param %q\n", name, c.Value(paramKey{}))

	select {
	case <-time.After(d):
		fmt.Printf("%s done \n", name)
	case <-c.Done():
		fmt.Printf("%s cancelled \n", name)
	}
}
