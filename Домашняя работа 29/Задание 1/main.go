package main

import (
	"fmt"
	"strconv"
)

func main() {
	var intdata string
	fmt.Println("Введите число или стово \"стоп\"")
	for {
		fmt.Scan(&intdata)
		if intdata == "стоп" {
			break
		}
		n,_ := strconv.Atoi(intdata)
		fc := squareNumber(n)
		productNumber(fc)

	}
}

func squareNumber(n int) chan int {
	secondchan := make(chan int)
	var result int
	go func() {
		result = n * n
		fmt.Println("Квадрат",result)
		secondchan <- result
	}()

	return secondchan
}
func productNumber(n chan int)  {
	go func() {
		result := <-n
		result = result * 2
		fmt.Println("Произведение",result)
	}()

}
