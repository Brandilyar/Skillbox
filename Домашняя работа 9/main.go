package main

import (
	"fmt"
	"math"
)

const (
	maxUint8  = math.MaxUint8
	maxUint16 = math.MaxUint16
	maxUint32 = math.MaxUint32
)

func main() {
	task_2()
}

func task_1() {
	var overflowUint8 uint8 = 0
	var overflowUint16 uint16 = 0
	var countoverflowuint8,countoverflowuint16 int = 0,0
	for i := 0; i <= maxUint32; i++ {
		overflowUint8 ++
		overflowUint16 ++
		if overflowUint8 == 0{
			countoverflowuint8++
		}
		if overflowUint16 == 0{
			countoverflowuint16++
		}
	}
	fmt.Printf("Количество переполнений для Uint8 %d, количество переполнений для Uint16 %d",countoverflowuint8,countoverflowuint16)
}
func task_2(){
	var numberone,numbertwo int16
	fmt.Println("Введите 2 числа")
	fmt.Scan(&numberone,&numbertwo)
	result := int32(numberone)*int32(numbertwo)
	switch {
	case result > 0:
		switch {
		case result <= math.MaxUint8:
			fmt.Println("uint8")
		case result <= math.MaxUint16:
			fmt.Println("uint16")
		default:
			fmt.Println("uint32")
		}
	case result < 0:
		switch {
		case result >= math.MinInt8:
			fmt.Println("int8")
		case result >= math.MinInt16:
			fmt.Println("int16")
		default:
			fmt.Println("int32")
		}
	
	}
}
