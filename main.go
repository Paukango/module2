package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// первая
	var strng string = "104"
	var num int = 35
	job1, err := strconv.Atoi(strng)
	if err != nil {
		fmt.Printf("Кал говн")
	}
	job2 := strconv.Itoa(num)
	fmt.Printf("%T \n", job1)
	fmt.Printf("%T \n", job2)
	fmt.Print("------------------- \n")
	// вторая
	var a float64
	var b float64
	fmt.Print("Введите скорость в м/сек для преобразования в км/ч \n")
	fmt.Scan(&a)
	fmt.Print("Введите скорость в м/с для преобразования в миль/ч \n")
	fmt.Scan(&b)
	a = a * 3.6
	AmericanVelocity := a
	b = b * 2.2369
	EuropeanVelocity := b
	fmt.Print("Скорость в км/ч - ", AmericanVelocity, "\n", "Скорость в миль/ч - ", math.Floor(EuropeanVelocity*100)/100)
	// писььs
}
