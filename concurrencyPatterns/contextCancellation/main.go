package main

import (
	"context"
	"fmt"
	"time"
)

func work(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("work done")
		default:
			fmt.Println("working...")
			time.Sleep(500 * time.Second)
		}
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go work(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	fmt.Println("work cancel as it is taking too much time")
	time.Sleep(time.Second)

}
