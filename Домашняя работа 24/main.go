package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	task_1()

	rand.Seed(time.Now().UnixNano())
	var slice []int
	for i := 0; i < 10; i++ {
		slice = append(slice, rand.Intn(100)+1)
	}
	fmt.Println(slice)
	func(slice []int) {
		for {
			check := 0
			for i := len(slice) - 1; i > 0; i-- {
				if slice[i-1] < slice[i] {
					slice[i], slice[i-1] = slice[i-1], slice[i]
					check++
				}
			}
			if check == 0 {
				break
			}
		}
		fmt.Println(slice)
	}(slice)
	fmt.Println(task_0_1("Привет", "приет"))
}

func task_1() {
	rand.Seed(time.Now().UnixNano())
	var slice []int
	for i := 0; i < 10; i++ {
		slice = append(slice, rand.Intn(100)+1)
	}
	fmt.Println(slice)
	for i := 1; i <= len(slice)-1; i++ {
		x := slice[i]
		j := i
		for j > 0 && slice[j-1] > x {
			slice[j] = slice[j-1]
			j--
		}
		slice[j] = x
	}
	fmt.Println(slice)
}
func task_0_1(stringOne, stringTwo string) bool {
	tempstring := []rune{}
	for _, v := range strings.ToLower(stringTwo) {
		if strings.Contains(strings.ToLower(stringOne), string(v)) {
			tempstring = append(tempstring, v)
		}

	}
	if string(tempstring) == stringTwo {
		return true
	} else {
		return false
	}

}
