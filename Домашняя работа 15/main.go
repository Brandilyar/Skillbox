package main

import "fmt"

func main() {
	doPanic(5)
}
func task_1() {
	var number int
	fmt.Println("Введите количетсво целых чисел")
	fmt.Scan(&number)
	numbers := make([]int, number)
	for i := 0; i < number; i++ {
		fmt.Println("Введите число", i+1)
		fmt.Scan(&numbers[i])
	}
	evensize := 0
	oddsize := 0
	for i := range numbers {
		if numbers[i]%2 == 0 {
			evensize++
		} else {
			oddsize++
		}
	}
	fmt.Printf("Количество четных чисел %v, количество нечетных чисел %v", evensize, oddsize)
}
func task_2() {
	var number int
	fmt.Println("Введите количетсво целых чисел")
	fmt.Scan(&number)
	numbers := make([]int, number)
	for i := 0; i < number; i++ {
		fmt.Println("Введите число", i+1)
		fmt.Scan(&numbers[i])
	}
	newnumbers := change(numbers)
	fmt.Println(newnumbers)
}
func change(numbers []int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
func doPanic(a int){
	if a > 1{
		doPanic(a-1)
	}
	panic(a)
}