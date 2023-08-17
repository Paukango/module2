package main

import "fmt"

func main() {

	readers := map[string]map[string][]string{

		"Петя": {
			"Книги":        []string{"Простаквашино", "Буратино"},
			"Пер. издания": []string{"Ну погоди", "Золушка", "Доктор Айболит "},
		},
		"Вова": {
			"Книги":        []string{"Приключения капитана Врунгеля", "Старик Хоттабыч"},
			"Пер. издания": []string{"Аленький цветочек", "Черная курица", "Бобовое зернышко"},
		},
		"Леша": {
			"Книги":        []string{"Волк и козлята", "Журавль и цапля"},
			"Пер. издания": []string{"Зайкина избушка"},
		},
		"Сережа": {
			"Книги":        []string{},
			"Пер. издания": []string{},
		},
	}

	booksCount, ok := (countReadersWithBooks(readers))
	fmt.Println("Общее кол-во изданий на руках у каждого читателя - ", booksCount, ok)

	EvenReader, ok := (countBooksEvenReader(readers))
	fmt.Println("Кол-во читателей с изданиями на руках -", EvenReader, ok)
}

func countReadersWithBooks(readers map[string]map[string][]string) (count int, err error) {

	var g int
	for _, v := range readers {
		for _, y := range v {
			if len(y) != 0 {
				g = len(y) + g
			}
		}
	}
	count = g
	return count, nil
}

func countBooksEvenReader(readers map[string]map[string][]string) (count int, err error) {

	var g int
	var x int
	for _, v := range readers {
		for _, y := range v {
			if len(y) != 0 {
				g++
			}
		}
		if g != 0 {
			x++
			g = 0
		}
	}
	count = x
	return count, nil
}
