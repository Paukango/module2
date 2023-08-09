package main

import "fmt"

func main() {
	jobDays := []string{"Пн", "Вт", "Ср", "Чт", "Пт"}
	weekends := []string{"Сб", "Вс"}
	week := make([]string, 0, 7)
	week = append(week, jobDays...)
	week = append(week, weekends...)
	fmt.Println(week)
}
