package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "my-key", "test value")
	viewContext(ctx)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go myProcess(ctx)
	<-ctx.Done()
	fmt.Println("We've exceeded the deadline")
	fmt.Println(ctx.Err())

}

func viewContext(ctx context.Context) {
	fmt.Printf("my value is %s\n", ctx.Value("my-key"))
}

func myProcess(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ohh, time out")
			return
		default:
			fmt.Println("Doing some process")
		}
		time.Sleep(time.Millisecond * 500)
	}
}
