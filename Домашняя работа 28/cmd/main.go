package main

import (
	"fmt"
	"io"
	"log"
	"student/pkg/student"
	studentstorage "student/pkg/student_storage"
)

func main() {
	var name string
	var age int
	var grade int
	stotage := studentstorage.NewStudentStorage()
	for {
		fmt.Println("Введите данные студента")
		_, err := fmt.Scan(&name, &age, &grade)
		stotage.Put(student.NewStudent(name, age, grade))
		fmt.Println(stotage)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
	for _, oneStudent := range stotage {
		fmt.Println(oneStudent.Name, oneStudent.Age, oneStudent.Grade)
	}
}
