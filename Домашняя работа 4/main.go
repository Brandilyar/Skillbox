package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	task_5()
}

//Предыдущее ДЗ
func task_0() {
	var str_one, str_two string
	fmt.Println("Введите 2 строки")
	fmt.Scan(&str_one, &str_two)
	if strings.Contains(str_one, str_two) {
		fmt.Println("Строка 2 содержится в строке 1")
	} else if strings.Contains(str_two, str_one) {
		fmt.Println("Строка 1 содержится в строке 2")
	} else {
		fmt.Println("Ни одна из строк не содержится в другой")
	}
}

// Баллы ЕГЭ
func task_1() {
	var resultExam int
	fmt.Println("Введите сумму баллов по 3 экзаменам")
	fmt.Scan(&resultExam)
	if resultExam < 0 {
		fmt.Println("Сумма баллов не может быть отрицательной")
	}
	if resultExam < 275 {
		fmt.Println("К сожалению вы не поступили")
	} else {
		fmt.Printf("Поздравляем! %d баллов достаточно для поступления в ВУЗ", resultExam)
	}
}

// Три числа
func task_2() {
	var number_1, number_2, number_3 int
	fmt.Println("Введите 3 числа")
	fmt.Scan(&number_1, &number_2, &number_3)
	if number_1 > 5 {
		fmt.Printf("Число %d больше 5", number_1)
	} else if number_2 > 5 {
		fmt.Printf("Число %d больше 5", number_2)
	} else if number_3 > 5 {
		fmt.Printf("Число %d больше 5", number_3)
	} else {
		fmt.Println("Среди введеных чисел нет того,которое больше 5")
	}

}

//  Банкомат
func task_3() {
	var summa int
	fmt.Println("Укажите сумму для снятия")
	fmt.Scan(&summa)
	if summa > 100000 {
		fmt.Printf("Извините. Указанная сумма в размере %d слишком большая. Укажите сумму в пределах от 100 до 100 000", summa)
	} else if summa < 100 {
		fmt.Printf("Извините. Указанная сумма в размере %d слишком маленькая. Укажите сумму в пределах от 100 до 100 000", summa)
	} else {
		fmt.Printf("Вы сняли сумму %d", summa)
	}
}

// Три числа: еще попытка
func task_4() {
	var number_1, number_2, number_3, countNumber int
	var slice []int
	fmt.Println("Введите 3 числа")
	fmt.Scan(&number_1, &number_2, &number_3)
	slice = append(slice, number_1, number_2, number_3)
	for _, value := range slice {
		if value >= 5 {
			countNumber++
		}
	}
	fmt.Printf("Количество чисел больше или равных 5 равно %d", countNumber)
}

// Ресторан
func task_5() {
	var dayWeek string
	var numberVizitor int
	var sum float32

	fmt.Println("Введите день недели")
	fmt.Scan(&dayWeek)
	fmt.Println("Введите количество гостей")
	fmt.Scan(&numberVizitor)
	fmt.Println("Введите сумму чека")
	fmt.Scan(&sum)
	dayWeek = strings.ToLower(dayWeek)
	dayWeek = strings.Title(dayWeek)
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
func task_6() {
	var n, k, number int
	var sliceStudent, sliceGroup []int
	fmt.Println("Введите количество студентов")
	fmt.Scan(&n)
	fmt.Println("Введите количество групп")
	fmt.Scan(&k)
	fmt.Println("Введите номер студента")
	fmt.Scan(&number)
	for i := 1; i <= n; i++ {
		sliceStudent = append(sliceStudent, i)
	}
	for i := 1; i <= k; i++ {
		sliceGroup = append(sliceGroup, i)
	}
	rand.Seed(time.Now().UnixNano())
	groupStudent := sliceGroup[rand.Intn(len(sliceGroup))]
	fmt.Printf("Студент %d попадет в группу %d", sliceStudent[number-1], groupStudent)
}
func task_7() {
	var sum float32
	fmt.Println("Введите сумму заработка")
	fmt.Scan(&sum)
	if sum >= 50000 {
		nalog := 0.3*(sum-50000) + 0.2*(50000-10000) + 0.13*10000
		fmt.Println(nalog)
	} else if 10000 < sum && sum < 50000 {
		nalog := 0.2*(sum-10000) + 0.13*10000
		fmt.Println(nalog)
	} else {
		nalog := 0.13 * sum
		fmt.Println(nalog)
	}
}
