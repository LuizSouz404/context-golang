package main

import (
	"context"
	"fmt"
	"time"
)

func numeros(v chan<- int) {
	for i := 0; i < 10; i++ {
		v <- i
		fmt.Printf("número %d escrito no channel \n", i)
	}
	close(v)
}

func main() {
	// ctx, cf := context.WithCancel(context.Background())
	ctx, cf := context.WithTimeout(context.Background(), time.Second*2)
	defer cf()

	// go func() {
	// 	time.Sleep(time.Second * 5)
	// 	cf()
	// 	fmt.Println("Timeout!")
	// }()

	c := make(chan int, 3)
	go numeros(c)

loopNum:
	for {
		select {
		case <-ctx.Done():
			break loopNum
		case v, ok := <-c:
			if ok {
				fmt.Printf("número %d lido no channel \n", v)
				time.Sleep(2 * time.Second)
			}
		}
	}
}
