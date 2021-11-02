package main

import (
	"fmt"
	"math"
	"time"
	"math/rand"
)

func main() {
	task_2()
}
func task_0_1() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Введите количесвто ip адресов")
	var x int
	fmt.Scan(&x)
	for i:=0;i<x;i++{
		fmt.Printf("%v.%v.%v.%v\n",rand.Intn(254) + 1,rand.Intn(254) + 1,rand.Intn(254) + 1,rand.Intn(254) + 1)
	}

}
func task_1() {
	var epsilon float64
	fmt.Println("Введите X для которого необходимо расчитать e^x")
	var x float64
	fmt.Scan(&x)
	fmt.Println("Введите с точностью до какого знака после запятой необходимо вычислить значение функции")
	var accuracy int
	fmt.Scan(&accuracy)
	epsilon = 1 / math.Pow(10, float64(accuracy))
	result := 0.0
	prev := 0.0
	fact := 1
	n := 0
	for {
		if n > 0 {
			fact *= n
		}
		result += math.Pow(x, float64(n)) / float64(fact)
		if math.Abs(result-prev) < epsilon {
			fmt.Println(result)
			break
		}
		n++
		prev = result
	}
}
func task_2() {
	var sum float64
	var p float64
	var year int
	fmt.Println("Введите сумму,которую вы вносите в банк")
	fmt.Scan(&sum)
	fmt.Println("Введите процент ежемесячной капитализации")
	fmt.Scan(&p)
	fmt.Println("Введите количество лет,на которое будет внесен вклад")
	fmt.Scan(&year)
	month := year * 12
	bankprofit := 0.0
	prevsum := 0.0
	for i := 1; i <= month; i++ {
		capitalization := sum * (p / 100)
		prevsum = sum + capitalization
		sum = math.Trunc(prevsum*100) / 100
		bankprofit += prevsum - sum
	}
	fmt.Println(sum, bankprofit)
}
