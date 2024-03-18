package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	time.Sleep(2 * time.Second)
	t := time.Until(now)
	fmt.Println("时间差距", t.Seconds())

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	go handle(ctx, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}

	time.Sleep(1 * time.Second)

	c, can := context.WithCancel(context.Background())
	go handle(c, 1500*time.Second)
	can()
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}

	ctx := context.Background()
	process(ctx)

	ctx = context.WithValue(ctx, "traceId", "abcedef")
	process(ctx)

	time.Sleep(1 * time.Second)

}

func process(ctx context.Context) {
	traceId, ok := ctx.Value("traceId").(string)
	if ok {
		fmt.Printf("process over. trace_id=%s\n", traceId)
	} else {
		fmt.Printf("process over. no trace_id\n")
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}

	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
