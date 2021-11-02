package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	task_0_1(0)
	fmt.Println("Выполнено")
}

func task_1() {
	rand.Seed(time.Now().UnixNano())
	var a [10]int
	var num int
	for i := 0; i < 10; i++ {
		a[i] = rand.Intn(10)
	}
	check := 0
	fmt.Println("Введите число")
	fmt.Scan(&num)
	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		if a[i] == num {
			check = len(a[i+1:])
			break
		}
	}
	if check > 0 {
		fmt.Printf("Количество числе после введенного числа равно %d \n", check)
	} else {
		fmt.Println(0)
	}
}

func task_2() {
	rand.Seed(time.Now().UnixNano())
	a := [12]int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var num int

	fmt.Println("Введите число")
	fmt.Scan(&num)
	fmt.Println(a)
	fmt.Println(find(a, num))
}
func find(a [12]int, value int) (index int) {
	index = -1
	min := 0
	max := 11
	for max >= min {
		middle := (min + max) / 2
		if a[middle] > value {
			max = middle - 1
		} else if a[middle] < value {
			min = middle + 1
		} else if middle == 0 || value != a[middle-1] {
			index = middle
			return
		} else {
			max = middle - 1
		}
	}
	return

}

func task_0_1(n int){
	defer recoveryfunc()
	a := 1/n
	fmt.Println(a)
	panic("Паника")
}

func recoveryfunc(){
	recoveryMessage:=recover()
	if recoveryMessage != nil{
		fmt.Println(recoveryMessage)
	}
	fmt.Println("This is recovery function...")
}