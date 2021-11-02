package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	common := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	even, noteven := task_1(common)
	fmt.Println("Четные",even)
	fmt.Println("Нечетные",noteven)
	
	sentences := []string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := []rune{'H', 'E', 'L', 'S', 'l'}
	task_2(sentences, chars)
}

func task_1(common []int) (even []int, noteven []int) {
	for i := range common {
		if common[i]%2 == 0 {
			even = append(even, common[i])
		} else {
			noteven = append(noteven, common[i])
		}
	}
	return
}
func task_2(sentences []string, chars []rune) {
	var result [][]string
	for _, one_string := range sentences {
		words := strings.Fields(one_string)
		for i, one_rune := range words[len(words)-1] {
			for _, one_chars := range chars {
				if one_chars == one_rune {
					result = append(result,([]string{string(one_chars),"position", strconv.Itoa(i)}))
				}
			}
		}
	}
	fmt.Println(result)
}
