package main

import (
	"fmt"
	"sync"
)


func main() {
	// ch := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1) // ?+

	go func() {
		defer wg.Done()

		fmt.Println("hello")
		// ch <- struct{}{}
	}()

	wg.Wait()

	// <- ch
	fmt.Println("world")
}
