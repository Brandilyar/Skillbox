package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	byfchan := make(chan int, 3)
	goroutines := make(chan struct{}, 10)
	iw := 0
	for i := 0; i < 100; i++ {
		if len(byfchan) == 3 {
			wg.Add(1)
			goroutines <- struct{}{}
			iw++
			go some_func(&wg, byfchan, goroutines, iw)
		} else {
			byfchan <- i
		}

	}
	close(byfchan)
	wg.Wait()
	close(goroutines)
}
func some_func(wg *sync.WaitGroup, c chan int, goroutines chan struct{}, i int) {
	
	for {
		value, ok := <-c
		if !ok {
			wg.Done()
			return
		}
		fmt.Printf("горутина %d, данные из канала %d \n", i, value)
		<-goroutines
	}

}
