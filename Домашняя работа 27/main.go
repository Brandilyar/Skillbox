package main

import (
	"fmt"
	"io"
	"log"
)

type Student struct {
	name  string
	age   int
	grade int
}

func newStudent(name string, age int, grade int) *Student {
	return &Student{name, age, grade}
}

func main() {
	var name string
	var age int
	var grade int
	result := map[string]*Student{}
	for {
		fmt.Println("Введите данные студента")
		_, err := fmt.Scan(&name, &age, &grade)
		result[name] = newStudent(name, age, grade)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
	fmt.Println("Студенты из хранилища")
	for _, oneStudent := range result {
		fmt.Println(oneStudent.name, oneStudent.age, oneStudent.grade)
	}

}
