package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	task_7()
}
func task_1() {
	min := int(math.Pow(10, 5))
	max := int(math.Pow(10, 6) - 1)
	discharges := []int{100000, 10000, 1000, 100, 10, 1}
	countbilet := 0
	for i := min; i <= max; i++ {
		countdischarges := []int{}
		for _, item := range discharges {
			countdischarges = append(countdischarges, i/item%10)
		}
		if countdischarges[0] == countdischarges[5] && countdischarges[1] == countdischarges[4] && countdischarges[2] == countdischarges[3] {
			countbilet++
		}
	}
	fmt.Println(countbilet)
}
func task_2() {
	fmt.Println("Введите размер шахмотной доски")
	var count int
	fmt.Scan(&count)
	for row := 1; row <= count; row++ {
		for column := 1; column <= count; column++ {
			if (row+column)%2 == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
func task_3() {
	fmt.Println("Введите размер елочки")
	var count int
	fmt.Scan(&count)
	for row, i, rep := 1, count, 1; row <= count; row, i, rep = row+1, i-1, rep+2 {
		fmt.Print(strings.Repeat(" ", i-1))
		fmt.Print(strings.Repeat("*", rep))
		fmt.Println()
	}
}
func task_4() {
	min := int(math.Pow(10, 5))
	max := int(math.Pow(10, 6) - 1)
	discharges := []int{100000, 10000, 1000, 100, 10, 1}
	maxDistance := 0
	happyticketone := 0
	happytickettwo := 0
	for i := min; i <= max; i++ {
		countdischarges := []int{}
		for _, item := range discharges {
			countdischarges = append(countdischarges, i/item%10)
		}
		if countdischarges[0]+countdischarges[1]+countdischarges[2] == countdischarges[3]+countdischarges[4]+countdischarges[5] {
			if happyticketone == happytickettwo {
				happyticketone = i
			} else {
				happytickettwo = happyticketone
				happyticketone = i
				if int(math.Abs(float64(happytickettwo-happyticketone))) > maxDistance {
					maxDistance = int(math.Abs(float64(happytickettwo - happyticketone)))
				}
			}
		}
	}
	fmt.Println(maxDistance)
}
func task_5() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100) + 1
	var answer int
	fmt.Println("Попробуйте угадать загаданные число")
	for {
		fmt.Println("Введите число")
		fmt.Scan(&answer)
		if answer == n {
			fmt.Println("Число угадано")
			break
		} else if answer > n {
			fmt.Println("Загаданное число меньше")
		} else if answer < n {
			fmt.Println("Загаданное число больше")
		} else if 1 < answer && answer < 100 {
			fmt.Println("Необходимо указать число в диапозоне от 1 до 100")
		}
	}
}
func task_6() {
	rand.Seed(time.Now().UnixNano())
	var slice []int
	for i := 0; i < 10; i++ {
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
func task_7() {
	rand.Seed(time.Now().UnixNano())
	var slice []int
	for i := 0; i < 10; i++ {
		slice = append(slice, rand.Intn(100)+1)
	}
	fmt.Println(slice)
	for i := 1; i <= len(slice)-1; i++ {
		x := slice[i]
		j := i
		for j > 0 && slice[j-1] > x {
			slice[j] = slice[j-1]
			j--
		}
		slice[j] = x
	}
	fmt.Println(slice)
}
