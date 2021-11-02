package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	task_8()
}

func task_0_1() {
	var resultExam int
	fmt.Println("Введите сумму баллов по 3 экзаменам")
	fmt.Scan(&resultExam)
	if resultExam < 275 && resultExam > 0 {
		fmt.Println("К сожалению вы не поступили")
	} else if resultExam < 0 {
		fmt.Println("Сумма баллов не может быть отрицательной")
	} else {
		fmt.Printf("Поздравляем! %d баллов достаточно для поступления в ВУЗ", resultExam)
	}
}

// Ресторан
func task_0_2() {
	var dayWeek string
	var numberVizitor int
	var sum float32

	fmt.Println("Введите день недели")
	fmt.Scan(&dayWeek)
	fmt.Println("Введите количество гостей")
	fmt.Scan(&numberVizitor)
	fmt.Println("Введите сумму чека")
	fmt.Scan(&sum)
	dayWeek = strings.ToLower(dayWeek) // Приведение к нижнему регистру
	dayWeek = strings.Title(dayWeek)   // Первая буква заглавная
	if dayWeek == "Понедельник" && numberVizitor <= 5 {
		sum = sum - (sum * 0.1)
	} else if dayWeek == "Понедельник" && numberVizitor > 5 {
		sum = sum - (sum * 0.1) + (sum * 0.1)
	} else if dayWeek == "Пятница" && sum > 10000 && numberVizitor <= 5 {
		sum = sum - (sum * 0.05)
	} else if dayWeek == "Пятница" && sum > 10000 && numberVizitor > 5 {
		sum = sum - (sum * 0.05) + (sum * 0.1)
	} else if numberVizitor > 5 {
		sum = sum + (sum * 0.1)
	}
	fmt.Printf("Сумма к оплате %f", sum)
}

func task_1() {
	fmt.Println("Введите значения координат X и Y")
	var x, y int
	fmt.Scan(&x, &y)
	if x > 0 && y > 0 {
		fmt.Println("Точка принадлежит первой четверти")
	} else if x < 0 && y > 0 {
		fmt.Println("Точка принадлежит второй четверти")
	} else if x < 0 && y < 0 {
		fmt.Println("Точка принадлежит трерьей четверти")
	} else if x > 0 && y < 0 {
		fmt.Println("Точка принадлежит четвертой четверти")
	}
}
func task_2() {
	fmt.Println("Введите 3 числа")
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	if x > 0 || y > 0 || z > 0 {
		fmt.Println("Хотя бы одно из введеных чисел положительно")
	} else {
		fmt.Println("Нет положительных чисел")
	}
}
func task_3() {
	fmt.Println("Введите 3 числа")
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	if x == y || y == z || z == x {
		fmt.Println("Вы ввели 2 или более одинаковых цисел")
	} else {
		fmt.Println("Нет одинаковых чисел")
	}
}
func task_4() {
	fmt.Println("Введите строимость товара")
	var x, y, z, sum int
	fmt.Scan(&sum)
	fmt.Println("Введите номиналы трех монет")
	fmt.Scan(&x, &y, &z)
	if x+y+z == sum || x+y == sum || x+z == sum || z+y == sum || x == sum || y == sum || z == sum {
		fmt.Println("Вы можете заплатить за товар без сдачи")
	} else {
		fmt.Println("Вы не можете заплатить за товар без сдачи")
	}
}
func task_5() {
	fmt.Println("Введите 3 ставки по депозитам")
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	if y < z && x < y {
		fmt.Printf("Ставки %d и %d являются максимальными", z, y)
	} else if x < z && x > y {
		fmt.Printf("Ставки %d и %d являются максимальными", z, x)
	} else {
		fmt.Printf("Ставки %d и %d являются максимальными", x, y)
	}
}
func task_6() {
	fmt.Println("Введите 3 числа")
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	d := b*b - 4*a*c
	if d > 0 {
		x1 := (float64(b*-1) + math.Sqrt(float64(d))) / float64(2*a)
		x2 := (float64(b*-1) - math.Sqrt(float64(d))) / float64(2*a)
		fmt.Printf("Корни уравнения %f и %f", x1, x2)
	} else if d == 0 {
		x1 := float64(b*-1) / float64(2*a)
		fmt.Printf("Корень уравнения %f", x1)
	} else {
		fmt.Println("Уравнение не имеет корней")
	}
}
func task_7() {
	fmt.Println("Введите 4-х значиный номер билет")
	var number int
	fmt.Scan(&number)
	number_1 := number / 1000
	number_2 := number / 100 % 10
	number_3 := number / 10 % 10
	number_4 := number % 10
	if number_1+number_2 == number_3+number_4 && strconv.Itoa(number_1)+strconv.Itoa(number_2) == strconv.Itoa(number_4)+strconv.Itoa(number_3) {
		fmt.Println("У Вас счасливый,зеркальный билет")
	} else if number_1+number_2 == number_3+number_4 {
		fmt.Println("У Вас счасливый")
	} else {
		fmt.Println("У Вас обычный билет")
	}

}
func task_8() {
	fmt.Println("Введите число от 1 до 10")
	rand.Seed(time.Now().UnixNano())
	var number int
	fmt.Scan(&number)
	for i := 0; i < 4; i++ {
		temp := rand.Intn(10)
		fmt.Printf("Ваше число %d? yes/no \n", temp)
		var ans string
		fmt.Scan(&ans)
		if strings.ToLower(ans) == "yes" && temp == number {
			fmt.Println("Число угадано")
			break
		} else if strings.ToLower(ans) == "yes" && temp != number {
			fmt.Println("Не обманывайте)))")
		} else if strings.ToLower(ans) == "no" && temp == number {
			fmt.Println("Не обманывайте))). Число угадано")
		} else if i == 3 {
			fmt.Println("Число не угадано")
		}

	}
}
