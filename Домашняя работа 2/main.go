package main

import (
	"fmt"
)

func main() {
	task_1()
	fmt.Println("///////////////////////////////")
	task_2()
	fmt.Println("///////////////////////////////")
	task_3()
	fmt.Println("///////////////////////////////")
	task_4()
	fmt.Println("///////////////////////////////")
}

func task_1(){
	fmt.Println("Task_1")
	lap := 4
	engine := 254
	wheels := 93
	steeringWheel := 49
	wind := 21
	rain := 17
	speed := engine+wheels+steeringWheel-wind-rain

	

	fmt.Println("===================")
	fmt.Print("Супер гонки. Круг ", lap, "\n")
	fmt.Println("===================")
	fmt.Print("Шумахер (", speed ,")\n")
	fmt.Println("===================")
	fmt.Println("Водитель: Шумахер")
	fmt.Print("Скорость: ", speed, "\n")
	fmt.Println("-------------------")
	fmt.Println("Оснащение")
	fmt.Print("Двигатель: +", engine, "\n")
	fmt.Print("Колеса: +", wheels, "\n")
	fmt.Print("Руль: +", steeringWheel, "\n")
	fmt.Println("-------------------")
	fmt.Println("Действия плохой погоды")
	fmt.Print("Ветер: -", wind, "\n")
	fmt.Print("Дождь: -", rain, "\n")
}
func task_2(){
	fmt.Println("Task_2")
	costProduct := 6400
	costDelivery := 350
	discount := 700
	fmt.Println("Стоимость товара",costProduct)
	fmt.Println("Стоимость доставки",costDelivery)
	fmt.Println("Размер скидки",discount)
	fmt.Println("===================")
	fmt.Println("Итого",costProduct+costDelivery-discount)
}
func task_3(){
	fmt.Println("Task_3")
	minutes :=480
	orderMinutes :=2
	cashierMinutes :=4
	fmt.Println("Эта программа рассчитает, сколько клиентов успеет обслужить кассир за смену.")
	fmt.Println("Введите длительность смены в минутах:",minutes)
	fmt.Println("Сколько минут клиент делает заказ? ",orderMinutes)
	fmt.Println("Сколько минут кассир собирает заказ? ",cashierMinutes)
	fmt.Println("-----Считаем-----")
	fmt.Println("За смену длиной", minutes, "минут кассир успеет обслужить", minutes/(orderMinutes+cashierMinutes), "клиентов.")
}
func task_4(){
	fmt.Println("Task_4")
	commonSum := 4000000
	entrances := 10
	flats := 40
	fmt.Println("Сумма, указанная в квитанции:",commonSum)
	fmt.Println("Подъездов в доме: ",entrances)
	fmt.Println("Квартир в каждом подъезде: ",flats)
	fmt.Println("----Результат-----")
	fmt.Println("Каждая квартира должна заплатить по", commonSum/(entrances*flats), "руб.")
}