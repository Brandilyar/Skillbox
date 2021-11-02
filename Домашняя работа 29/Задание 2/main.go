package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	exitFunc(c)


}

func exitFunc(c chan os.Signal) {
	for i:=0;;i++ {
		select {
		case v := <-c:
			if v == os.Interrupt {
				close(c)
				fmt.Println("Выхожу из программы")
				return
			}
		default:
			fmt.Println(i*i)
		}
	}
}
