package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ctx, cancle := context.WithCancel(context.Background())
	ch := make(chan int, 100)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for {
				select {
				case c := <-ch:
					fmt.Println(c)
				case <-ctx.Done():
					fmt.Println("exit")
					wg.Done()
					return
				}
			}
		}()
	}
	for i := 1; i < 101; i++ {
		ch <- i
	}
	cancle()
	wg.Wait()
}
