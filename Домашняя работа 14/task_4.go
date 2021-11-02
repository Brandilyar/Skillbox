package main

import (
	"fmt"
	"math/rand"
	"time"
)

var c = 100

func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	n = f1(n)
	n = f2(n)
	n = f3(n)
	fmt.Println(n)
}
func f1(n int) int {
	n = n + c
	return n
}
func f2(n int) int {
	n = n + c
	return n
}
func f3(n int) int {
	n = n + c
	return n
}