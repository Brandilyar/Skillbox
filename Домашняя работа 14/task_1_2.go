package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	task_2()
}

func genPointer(a, b int) (x, y int) {
	x = rand.Intn(a)
	y = rand.Intn(b)
	return
}
func changePointer(a, b int) (x, y int) {
	x = 2*a + 10
	y = -3*b - 5
	return
}

func task_1(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

func task_2() {
	x1, y1 := genPointer(10, 20)
	x2, y2 := genPointer(10, 20)
	x3, y3 := genPointer(10, 20)

	newx1, newy1 := changePointer(x1, y1)
	newx2, newy2 := changePointer(x2, y2)
	newx3, newy3 := changePointer(x3, y3)
	fmt.Println(newx1, newy1)
	fmt.Println(newx2, newy2)
	fmt.Println(newx3, newy3)
}
