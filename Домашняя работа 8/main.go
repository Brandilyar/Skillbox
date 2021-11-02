package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	task_3()
}
func task_0_1() {
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
func task_0_2() {
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
func task_1() {
	fmt.Println("Введите название месяца")
	var month string
	fmt.Scan(&month)
	switch month {
	case "декабрь", "январь", "февраль":
		fmt.Println("Время года зима")
	case "март", "апрель", "май":
		fmt.Println("Время года весна")
	case "июнь", "июль", "август":
		fmt.Println("Время года лето")
	case "сентябрь", "октябрь", "ноябрь":
		fmt.Println("Время года лето")
	default:
		fmt.Println("Неверное название месяца")
	}
}
func task_2() {
	fmt.Println("Введите название буднего дня недели (пн,вт,ср,чт,пт)")
	var day string
	fmt.Scan(&day)
	switch day {
	case "пн":
		fmt.Println("Понедельник, Вторник, Среда, Четверг, Пятница")
	case "вт":
		fmt.Println("Вторник, Среда, Четверг, Пятница")
	case "ср":
		fmt.Println("Среда, Четверг, Пятница")
	case "чт":
		fmt.Println("Четверг, Пятница")
	case "пт":
		fmt.Println("Пятница")
	default:
		fmt.Println("Неверное название дня недели")
	}
}
func task_3() {
	fmt.Println("Введите количество поситителей")
	var count int
	var bills []int
	fmt.Scan(&count)
	for i := 1; i <= count; i++ {
		var bill int
		fmt.Printf("Введите купюру %d клиента (5, 10, 20)\n", i)
		fmt.Scan(&bill)
		if bill == 5 || bill == 10 || bill == 20 {
			bills = append(bills, bill)
		} else {
			fmt.Println("Введена неверная купюра")
			i--
		}

	}
	fmt.Println(lemonadeChange(bills))
}
func lemonadeChange(bills []int) bool {
	check := false
	for i := 0; i < len(bills)-1; i++ {
		switch bills[i] {
		case 5:
			check = true
		case 10:		
			for j := 0; j < len(bills)-1; j++ {
				switch bills[j] {
				case 5:
					bills = append(bills[:j], bills[j+1:]...)
					check = true
				default:
					check = false
				}
			}
		case 20:
			checksum := 0
			for j := 0; j < len(bills)-1; j++ {
				checksum = checksum + bills[j]
				switch {
				case checksum == 15:
					check = true
				case checksum < 15:
					bills = append(bills[:j],bills[j+1:]...)
					check = false
				case checksum > 15:
					checksum = checksum - bills[j]
					check = false
				default:				
					check = false
				}
				
			}
		}
	}
	return check
}
