package main

import "fmt"

func main() {

	var i int
	reader := map[string]map[string][]string{

		"Читатель 1": {
			"Книга 1": []string{"Издание 2022"},
			"Книга 2": []string{"Издание 2021"},
			"Книга 3": []string{"Издание 2020"},
			"Книга 4": []string{"Издание 2020"},
		},
		"Читатель 2": {
			"Книга 1": []string{"Издание 2023"},
			"Книга 2": []string{"Издание 2020"},
			"Книга 3": []string{"Издание 2029"},
		},
		"Читатель 3": {
			"Книга 1": []string{"Издание 2011"},
			"Книга 2": []string{"Издание 2012"},
			"Книга 3": []string{"Издание 2013"},
		},
	}

	fmt.Println("Всего читателей - ", len(reader))
	for range reader["Читатель 1"] {
		i++
	}

	for range reader["Читатель 2"] {
		i++
	}

	for range reader["Читатель 3"] {
		i++

	}
	fmt.Println("Количество изданий на руках - ", i)
}
