package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	err := executeTaskWithTimeout(3)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	} else {
		fmt.Println("Long function is finished")
	}
	// runtime.Goexit()
}

func executeTaskWithTimeout(timeout int) error {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancelCtx()

	doneCh := make(chan int)
	go func() {
		executeTask()
		// doneCh <- 1
		close(doneCh)
	}()

	select {
	case <-doneCh:
		return nil
	case <-ctx.Done():
		return ctx.Err()
		break
	}

	return nil
}

func executeTask() {
	time.Sleep(10 * time.Second)
	fmt.Println("done task")
}
