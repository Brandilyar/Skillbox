package main

import (
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	task_1()
}

func task_0_1() {
	var one, two, result big.Int
	one.Exp(big.NewInt(10), big.NewInt(50), nil)
	two.Exp(big.NewInt(10), big.NewInt(60), nil)
	result.Mul(&one, &two)
	fmt.Println(&result)

}

func task_1() {
	s := "Go is an Open soUrce ProgramMing Language that makes it Easy to build simple, reliable, and efficient Software"
	countword := 0
	for strings.Contains(s," "){
		indexSpace := strings.Index(s," ")
		word := s[:indexSpace]
		b := regexp.MustCompile(`[A-Z]`).FindAll([]byte(word[:1]), -1)
		for _,v := range b {
			if len(v) > 0{
				countword++
			}
		}
		s = s[indexSpace+1:]
	}
	b := regexp.MustCompile(`[A-Z]`).FindAll([]byte(s), -1)
	for _,v := range b {
		if len(v) > 0{
			countword++
		}
	}
	fmt.Printf("Количество слов начинающихся с большой буквы равно %d",countword)
}
func task_2() {
	s := "a10 10 20b 20 30c30 30 dd"
	sliceString := strings.Split(s, " ")
	var result []int64
	for _, v := range sliceString {
		i, err := strconv.ParseInt(v, 16, 0)
		if err != nil {
			continue
		} else {
			result = append(result, i)
		}
	}
	fmt.Println(result)
}
