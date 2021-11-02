package main

import (
	"fmt"
	"math"
)

func main() {
	task_5()
}

//ДЗ № 5
func task_0_1() {

	fmt.Println("Введите значения координат X и Y")
	var x, y int
	fmt.Scan(&x, &y)
	switch {
	case x > 0 && y > 0:
		fmt.Println("Точка принадлежит первой четверти")
	case x < 0 && y > 0:
		fmt.Println("Точка принадлежит второй четверти")
	case x < 0 && y < 0:
		fmt.Println("Точка принадлежит трерьей четверти")
	case x > 0 && y < 0:
		fmt.Println("Точка принадлежит четвертой четверти")
	}
}

//ДЗ № 5
func task_0_4() {
	fmt.Println("Введите строимость товара")
	var x, y, z, sum int
	fmt.Scan(&sum)
	fmt.Println("Введите номиналы трех монет")
	fmt.Scan(&x, &y, &z)
	switch {
	case x+y+z == sum || x+y == sum || x+z == sum || z+y == sum || x == sum || y == sum || z == sum:
		fmt.Println("Вы можете заплатить за товар без сдачи")
	default:
		fmt.Println("Вы не можете заплатить за товар без сдачи")
	}
}

//ДЗ № 5
func task_0_5() {
	fmt.Println("Введите 3 ставки по депозитам")
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	switch {
	case y < z && x < y:
		fmt.Printf("Ставки %d и %d являются максимальными", z, y)
	case x < z && x > y:
		fmt.Printf("Ставки %d и %d являются максимальными", z, x)
	default:
		fmt.Printf("Ставки %d и %d являются максимальными", x, y)
	}
}

//ДЗ № 5
func task_0_6() {
	fmt.Println("Введите 3 числа")
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	d := b*b - 4*a*c
	switch {
	case d > 0:
		x1 := (float64(b*-1) + math.Sqrt(float64(d))) / float64(2*a)
		x2 := (float64(b*-1) - math.Sqrt(float64(d))) / float64(2*a)
		fmt.Printf("Корни уравнения %f и %f", x1, x2)
	case d == 0:
		x1 := float64(b*-1) / float64(2*a)
		fmt.Printf("Корень уравнения %f", x1)
	default:
		fmt.Println("Уравнение не имеет корней")
	}
}

//ДЗ № 6
func task_1() {
	fmt.Println("Введите число")
	var a int
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Println(i)
	}
}

//ДЗ № 6
func task_2() {
	fmt.Println("Введите 2 числа")
	var a, b int
	fmt.Scan(&a, &b)
	sum := a + b
	if sum > a {
		for a != sum {
			a++
		}
	} else {
		for a != sum {
			a--
		}
	}
	fmt.Println("Первое число достигло суммы первого и второго числа", sum)

}

//ДЗ № 6
func task_3() {
	fmt.Println("Введите цену товара и скидку в процентах")
	var price, discount float64
	fmt.Scan(&price, &discount)
	result := price * (discount / 100)
	switch {
	case discount > 30:
		fmt.Println("Скидка не может быть больше 30%")
	case result <= 2000 && discount <= 30:
		fmt.Printf("Скидка составила %f", float64(result))
	default:
		fmt.Println("Скидка составила 2000 руб.")
	}
}

//ДЗ № 6
func task_4() {
	a, b, c := 0, 0, 0
	for {
		if a == 10 && b == 100 && c == 1000 {
			break
		}
		if a != 10 {
			a++
		}
		if b != 100 {
			b++
		}
		if c != 1000 {
			c++
		}
	}
	fmt.Println(a, b, c)
}

//ДЗ № 6
func task_5() {
	fmt.Println("Введите емкость 3-х корзин")
	var sizeOne, sizeTwo, sizeThree int
	fmt.Scan(&sizeOne, &sizeTwo, &sizeThree)
	basketOne, basketTwo, basketThree := 0, 0, 0
	baskets := []*int{&basketOne,&basketTwo, &basketThree}
	sizeslice := []int{sizeOne,sizeTwo,sizeThree}
	for i:=0;basketOne < sizeOne || basketTwo < sizeTwo || basketThree < sizeThree;i++ {
		if i == 3{
			i = 0
		}
		if *baskets[i] < sizeslice[i] {
			*baskets[i]++
		}
	}
	fmt.Printf("В первой корзине %d яблок, во второй корзине %d яблок, в трерьей корзине %d яблок", basketOne, basketTwo, basketThree)
}

//ДЗ № 6
func task_6() {
	const (
		minFloor     = 1
		maxFloor     = 24
		maxPassenger = 2
		floorFour    = 4
		floorSeven   = 7
		floorTen     = 10
	)
	passengerInElevator := 0
	floorElevator := 1
	totalpassenger := 0
	direction := true
	passengerOnTenFloor := 3
	passengerOnSevenFloor := 3
	passengerOnFourFloor := 3
	for totalpassenger < 9 {
		if floorElevator == minFloor && direction {
			direction = false
			totalpassenger += passengerInElevator
			passengerInElevator = 0
			fmt.Printf("Перевезено %d пассажиров \n", totalpassenger)
			for floorElevator != maxFloor {
				floorElevator += 1
			}
		} else if floorElevator == maxFloor && !direction {
			direction = true
			for floorElevator != minFloor {
				floorElevator -= 1
				if floorElevator == floorTen && passengerInElevator != maxPassenger && passengerOnTenFloor != 0 {
					if passengerOnTenFloor > maxPassenger-passengerInElevator {
						passengerInElevator += maxPassenger - passengerInElevator
						passengerOnTenFloor -= maxPassenger - passengerInElevator
					} else {
						passengerInElevator += passengerOnTenFloor
						passengerOnTenFloor -= passengerOnTenFloor
					}
					fmt.Printf("На 10 этаже осталось %d пассажиров \n", passengerOnTenFloor)
					floorElevator -= 1
				} else if floorElevator == floorSeven && passengerInElevator != maxPassenger && passengerOnSevenFloor != 0 {
					enterPassenger := maxPassenger - passengerInElevator
					if passengerOnSevenFloor > enterPassenger {
						passengerInElevator += enterPassenger
						passengerOnSevenFloor -= enterPassenger
					} else {
						passengerInElevator += passengerOnSevenFloor
						passengerOnSevenFloor -= passengerOnSevenFloor
					}
					fmt.Printf("На 7 этаже осталось %d пассажиров \n", passengerOnSevenFloor)
					floorElevator -= 1
				} else if floorElevator == floorFour && passengerInElevator != maxPassenger && passengerOnFourFloor != 0 {
					enterPassenger := maxPassenger - passengerInElevator
					if passengerOnFourFloor > enterPassenger {
						passengerInElevator += enterPassenger
						passengerOnFourFloor -= enterPassenger
					} else {
						passengerInElevator += passengerOnFourFloor
						passengerOnFourFloor -= passengerOnFourFloor
					}
					fmt.Printf("На 4 этаже осталось %d пассажиров \n", passengerOnSevenFloor)
					floorElevator -= 1
				}
			}
		}
	}
}
