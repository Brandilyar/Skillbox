package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {

	var str, substr string
	flag.StringVar(&str, "str", "", "set str")
	flag.StringVar(&substr, "substr", "", "set substr")
	flag.Parse()
	fmt.Println(task_1(str, substr))
}
func task_1(stringOne, stringTwo string) bool {
	tempstring := []rune{}
	for _, v := range strings.ToLower(stringTwo) {
		for _, value := range strings.ToLower(stringOne){
			if value == v {
				tempstring = append(tempstring, v)
			}
		}
	}
	if string(tempstring) == stringTwo {
		return true
	} else {
		return false
	}
}
