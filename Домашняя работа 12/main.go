package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	task_1()
}
func task_1() {
	file, err := os.Create("log.txt")
	defer file.Close()
	if err := os.Chmod("log.txt", 0666); err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
	
	var inputstring string
	for i := 1; ; i++ {
		fmt.Println("Введите строку")
		fmt.Scan(&inputstring)
		if strings.ToLower(inputstring) == "выход" {
			break
		}
		file.WriteString(fmt.Sprint("№ строки ", i, " дата ", time.Now().Format("2006-01-02 15:04:05"), " ", inputstring, "\n"))
		file.Sync()
	}
}
func task_2() {
	file, err := os.Open("log.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	data, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}
	buf := make([]byte, data.Size())
	if _, err := io.ReadFull(file, buf); err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf))
}
func task_3() {
	var buf bytes.Buffer
	file, err := os.Create("log.txt")
	defer file.Close()
	if err := os.Chmod("log.txt", 0444); err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}	
	var inputstring string
	for i := 1; ; i++ {
		fmt.Println("Введите строку")
		fmt.Scan(&inputstring)
		if strings.ToLower(inputstring) == "выход" {
			break
		}
		buf.WriteString(fmt.Sprint("№ строки ", i, " дата ", time.Now().Format("2006-01-02 15:04:05"), " ", inputstring, "\n"))
		
	}
	if err := ioutil.WriteFile("log.txt", buf.Bytes(), 0); err != nil {
		panic(err)
	}
}
func task_4() {
	var buf bytes.Buffer
	file, err := os.Create("log.txt")
	if err := os.Chmod("log.txt", 0666); err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100) + 1
	fmt.Println("Введите число от 1 до 100")
	buf.WriteString("Введите число от 1 до 100 \n")
	for {
		var inputint int
		for {
			fmt.Scan(&inputint)
			buf.WriteString(fmt.Sprint("Введено число ", inputint, "\n"))
			if inputint < 0 || inputint > 100 {
				fmt.Println("Введите число от 1 до 100")
				buf.WriteString("Введите число от 1 до 100 \n")
			} else {
				break
			}
		}
		if inputint == n {
			fmt.Println("Ура! Вы угадали число")
			buf.WriteString("Ура! Вы угадали число \n")
			if err := ioutil.WriteFile("log.txt", buf.Bytes(), 0); err != nil {
				panic(err)
			}
			return
		} else if inputint < n {
			fmt.Println("Загаданное число больше")
			buf.WriteString("Загаданное число больше \n")
		} else {
			fmt.Println("Загаданное число меньше")
			buf.WriteString("Загаданное число меньше \n")
		}
	}
}


func task4_1(){
	f, err := os.Open("log.txt")
	defer f.Close()
	if err != nil{
		fmt.Println(err)
	}
	buf,err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf))
}

func task_5(n int, left int, right int, ans string,slice []string) []string {
	fmt.Println(ans)
	if left+right == 2*n {
		slice = append(slice,ans)
		return slice
	}
	if left < n {
		fmt.Println("left",left,"right",right)
		slice = task_5(n, left+1, right, ans+"(",slice)
	}
	if left > right {
		fmt.Println("right",right,"left",left)
		slice = task_5(n, left, right+1, ans+")",slice)
	}
	return slice
}

func task_6() {
	var input int
	fmt.Println("Введите число скобок")
	fmt.Scan(&input)
	ans := ""
	left := 0
	right := 0
	slice := []string{}
	slice = task_5(input, left, right, ans,slice)
	fmt.Println(slice)
}
