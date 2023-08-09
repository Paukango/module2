package main

import "fmt"

func main() {

	days := []string{"Пн", "Вт", "Ср", "Чт", "Пт", "Сб", "Вс"}
	jobDays := make([]string, 0, 5)
	jobDays = days[:5]
	days = days[5:7]
	fmt.Println(jobDays, days)

}
