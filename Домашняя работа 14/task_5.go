package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	resultfunc(f1, f2, f3, f4)
}

func resultfunc(rangefunc ...func(int)) {
	for i := range rangefunc {
		f := rangefunc[len(rangefunc)-1-i]
		f(rand.Intn(101))
	}
}
func f1(n int) {
	fmt.Println("Функция 1", n)
}
func f2(n int) {
	fmt.Println("Функция 2", n)
}
func f3(n int) {
	fmt.Println("Функция 3", n)
}
func f4(n int) {
	fmt.Println("Функция 4", n)
}
