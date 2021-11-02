package main

import "fmt"

func main() {
	result(one, two)
}

func one(one int) {
	fmt.Println("func one", one)
}
func two(two int) {
	fmt.Println("func two", two)
}
func result(one func(int), two func(int)) {
	two(2)
	one(1)
}
