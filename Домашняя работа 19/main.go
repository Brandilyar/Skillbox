package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	one := [4]int{2, 3, 4, 1}
	two := [5]int{6, 5, 8, 7, 9}
	fmt.Println(task_1(one, two))
	task_2()
}

//ДЗ 15 На слайсах
func task_0_1() {
	//ДЗ 15 На слайсах
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

//ДЗ 15 На слайсах
func task_0_2() {

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

// ДЗ 19
func task_1(one [4]int, two [5]int) []int {
	var result []int
	result = append(result, one[:]...)
	result = append(result, two[:]...)
	for {
		check := 0
		for i := 0; i < len(result)-1; i++ {
			if result[i+1] < result[i] {
				result[i], result[i+1] = result[i+1], result[i]
				check++
			}
		}
		if check == 0 {
			break
		}
	}
	return result
}

// ДЗ 19
func task_2() {
	rand.Seed(time.Now().UnixNano())
	var slice []int
	for i := 0; i < 6; i++ {
		slice = append(slice, rand.Intn(100)+1)
	}
	fmt.Println(slice)
	for {
		check := 0
		for i := 0; i < len(slice)-1; i++ {
			if slice[i+1] < slice[i] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
				check++
			}
		}
		if check == 0 {
			break
		}
	}
	fmt.Println(slice)
}
