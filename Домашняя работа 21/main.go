package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(task_1(1, 1, 1))
	task_2(1,2,func(x,y int)int{return x + y})
	task_2(1,2,func(x,y int)int{return x * y})
	task_2(1,2,func(x,y int)int{return x / y})
}

func task_1(x int16, y uint8, z float32) float32 {
	s := 2*float32(x) + float32(math.Pow(float64(y), 2)) - 3/z
	return s
}

func task_2(x,y int,A func(int,int)int){
	defer fmt.Println(A(x,y))
	fmt.Println("Вызов функции")
}