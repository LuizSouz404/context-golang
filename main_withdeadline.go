// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ctx, cancel := context.WithDeadline(
// 		context.Background(),
// 		time.Now().Add(10*time.Second),
// 	)
// 	fmt.Println(time.Now().Add(10 * time.Second).Format(time.RFC3339))
// 	defer cancel()

// 	printUntilCancel(ctx)
// }

// func printUntilCancel(ctx context.Context) {
// 	count := 0
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Cancel signal received, exiting...")
// 			return
// 		default:
// 			time.Sleep(1 * time.Second)
// 			fmt.Printf("Printing until cancel, number = %d \n", count)
// 			count++
// 		}
// 	}
// }
