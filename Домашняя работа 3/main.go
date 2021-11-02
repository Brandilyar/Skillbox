package main

import "fmt"

func main() {
	task_4_3()
}
// Конструктор квестов
func task_1(){	
	fmt.Println("Конструктор квестов")
	fmt.Println("Здравствуйте,введите название планеты")
	var planetName,nameStar,nameRanger string
	var money int
	fmt.Scan(&planetName)
	fmt.Println("Здравствуйте,введите название звездной системы")
	fmt.Scan(&nameStar)
	fmt.Println("Здравствуйте,введите имя рейнджера")
	fmt.Scan(&nameRanger)
	fmt.Println("Здравствуйте,введите количество вознаграждения")
	fmt.Scan(&money)
	fmt.Println("Как вам,", nameRanger, "известно, мы - раса мирная, поэтому на наших военных кораблях используются наемники с других планет.")
	fmt.Println("Система набора отработана давным-давно. Обычно мы приглашаем на наши корабли людей с планеты",planetName, "системы", nameStar,".")
	fmt.Println("Но случилась неприятность - в связи с большими потерями, в последнее время, престиж профессии сильно упал, и теперь не так-то просто завербовать призывников.")
	fmt.Println("Главный комиссар планеты", planetName,"впрочем, отлично справлялся, но недавно его наградили орденом Сутулого с закруткой на спине, и его на радостях парализовало! Призыв под угрозой срыва!")
	fmt.Println(nameRanger, "вы должны прибыть на планету", planetName,"системы", nameStar ,"и помочь выполнить план призыва. Мы готовы выплатить вам премию в", money,"кредитов за эту маленькую услугу.")
}
// Симулятор мартшрутки
func task_2(){
	var quantityPassengers,PassengersIn,PassengersOut,totalPassengers int
	var salary,taxes,fuel,repair,totalsum float32
	fmt.Println("Прибываем на остановку Улица Программистов. В салоне пассажиров:",quantityPassengers)
	fmt.Println("Сколько пассажиров вышло на остановке?")
	fmt.Scan(&PassengersOut)
	fmt.Println("Сколько пассажиров зашло на остановке?")
	fmt.Scan(&PassengersIn)
	totalPassengers += PassengersIn // Общее количество пассажиров воспользовавшихся маршруткой 
	quantityPassengers -= PassengersOut // Изменение количества пассажиров в маршрутке
	quantityPassengers += PassengersIn // Изменение количества пассажиров в маршрутке
	fmt.Println("Отправляемся с остановки Улица Программистов. В салоне пассажиров:",quantityPassengers)
	fmt.Println("-----------Едем---------")
	fmt.Println("Прибываем на остановку Проспект Алгоритмов. В салоне пассажиров:",quantityPassengers)
	fmt.Println("Сколько пассажиров вышло на остановке?")
	fmt.Scan(&PassengersOut)
	fmt.Println("Сколько пассажиров зашло на остановке?")
	fmt.Scan(&PassengersIn)
	totalPassengers += PassengersIn // Общее количество пассажиров воспользовавшихся маршруткой 
	quantityPassengers -= PassengersOut // Изменение количества пассажиров в маршрутке
	quantityPassengers += PassengersIn // Изменение количества пассажиров в маршрутке
	fmt.Println("Отправляемся с остановки Проспект Алгоритмов. В салоне пассажиров:",quantityPassengers)
	fmt.Println("-----------Едем---------")
	fmt.Println("Прибываем на остановку Проспект Боголюбова. В салоне пассажиров:",quantityPassengers)
	fmt.Println("Сколько пассажиров вышло на остановке?")
	fmt.Scan(&PassengersOut)
	fmt.Println("Сколько пассажиров зашло на остановке?")
	fmt.Scan(&PassengersIn)
	totalPassengers += PassengersIn // Общее количество пассажиров воспользовавшихся маршруткой 
	quantityPassengers -= PassengersOut // Изменение количества пассажиров в маршрутке
	quantityPassengers += PassengersIn // Изменение количества пассажиров в маршрутке
	fmt.Println("Отправляемся с остановки Проспект Боголюбова. В салоне пассажиров:",quantityPassengers)
	fmt.Println("-----------Едем---------")
	fmt.Println("Прибываем на остановку Улица Попова. В салоне пассажиров:",quantityPassengers)
	fmt.Println("Сколько пассажиров вышло на остановке?")
	fmt.Scan(&PassengersOut)
	fmt.Println("Сколько пассажиров зашло на остановке?")
	fmt.Scan(&PassengersIn)
	totalPassengers += PassengersIn // Общее количество пассажиров воспользовавшихся маршруткой 
	quantityPassengers -= PassengersOut // Изменение количества пассажиров в маршрутке
	quantityPassengers += PassengersIn // Изменение количества пассажиров в маршрутке
	fmt.Println("Отправляемся с остановки Улица Попова. В салоне пассажиров:",quantityPassengers)
	fmt.Println("Результат работы")
	totalsum = float32(totalPassengers) * 20
	salary = totalsum * 0.25
	fuel = totalsum * 0.2
	taxes = totalsum * 0.2
	repair = totalsum * 0.2
	fmt.Println("Всего заработали:",totalsum,"руб.")
	fmt.Println("Зарплата водителя:",salary,"руб.")
	fmt.Println("Расходы на топливо:",fuel,"руб.")
	fmt.Println("Налоги:",taxes,"руб.")
	fmt.Println("Расходы на ремонт машины:",repair,"руб.")
	fmt.Println("Итого заработали:",totalsum-salary-fuel-taxes-repair,"руб.")
}
// Перестановка значений 2 переменных с помощью 3 переменной
func task_3(){
	a := 42
	b := 153    
	fmt.Println("a:", a)   
	fmt.Println("b:", b)	
	c:= a
	a = b
	b = c
	fmt.Println("a:", a)  
	fmt.Println("b:", b)
}
// Расчет высоты бамбука на середину 3 дня
func task_4_1(){	
	height := 100 //Высота бамбука
	growth := 50 // Ежеденвный рост
	ate := 20 // Ежеденвное поедание бамбука
	height = (growth - ate)*2 + growth/2 + height
	fmt.Println("Высота бимбука на середину третьего дня:",height)
}
//Расчет требуемых дней роста для достижения указанной высоты
func task_4_2(){	
	height := 100 //Высота бамбука
	growth := 50 // Ежеденвный рост
	ate := 20 // Ежеденвное поедание бамбука
	day := 1 // Стартовое значение дня
	for ; height <= 300; day++ {
		height = (growth - ate) + height
	}
	fmt.Println("Понадобится",day-1,"дней чтобы бамбук достиг высоты 300 сантиметров")
}
//Расчет высоты бамбука в конкретный день с использованием вводимых пользователем данных
func task_4_3(){	
	var height,growth,ate,day int
	fmt.Println("Введите стартовую высоту саженца бамбука в сантиметрах")
	fmt.Scan(&height)
	fmt.Println("Введите ежедневный рост бамбука в сантиметрах")
	fmt.Scan(&growth)
	fmt.Println("Введите скорость поедания бамбука в сантиметрах")
	fmt.Scan(&ate)
	fmt.Println("Введите день на который вы хотите узнать высоту бамбука")
	fmt.Scan(&day)
	height = (growth - ate) * day + height
	fmt.Println("На",day,"день бамбук будет иметь высоту",height,"сантиметров")
}
//Расчет требуемых дней роста с использованием вводимых пользователем данных
func task_4_4(){	
	var height,growth,ate,resultHeight int
	fmt.Println("Введите стартовую высоту саженца бамбука в сантиметрах")
	fmt.Scan(&height)
	fmt.Println("Введите ежедневный рост бамбука в сантиметрах")
	fmt.Scan(&growth)
	fmt.Println("Введите скорость поедания бамбука в сантиметрах")
	fmt.Scan(&ate)
	fmt.Println("Введите высоту взрослого бамбука в сантиметрах")
	fmt.Scan(&resultHeight)
	day:=1
	for ; height <= resultHeight; day++ {
		height = (growth - ate) + height
	}
	fmt.Println("Понадобится",day-1,"дней чтобы бамбук достиг высоты ",resultHeight,"сантиметров")
}
//Перестановка значений переменных без использования 3 переменной
func task_5(){	
	one := 1
	two := 2
	fmt.Println("Переменная one",one,"Переменная two",two)
	one, two = two, one
	fmt.Println("Переменная one",one,"Переменная two",two)
}