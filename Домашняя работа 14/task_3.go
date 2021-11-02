package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(task(rand.Intn(10)))
}
func f1(n int) (n1 int) {
	n1 = n * rand.Intn(100)
	return
}
func f2(n int) (n1 int){
	n1 = n + rand.Intn(100)
	return
}
func task(n int) (n1 int){
	n1 = f1(n)
	n1 = f2(n1)
	return
}