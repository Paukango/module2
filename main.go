package main

import "fmt"

func main() {

	readers := map[string]map[string][]string{

		"Петя": {
			"Книги":        []string{"Жук", "Паук"},
			"Пер. издания": []string{"Старик", "Дед", "Прадед"},
		},
		"Вова": {
			"Книги":        []string{},
			"Пер. издания": []string{"Старик", "Дед", "Прадед"},
		},
		"Леша": {
			"Книги":        []string{"Дук", "Дудк"},
			"Пер. издания": []string{"пропо"},
		},
		"Сережа": {
			"Книги":        []string{},
			"Пер. издания": []string{},
		},
	}

	//fmt.Println(countReadersWithBooks(readers))

	fmt.Println(countBooksEvenReader(readers))

}

////func countReadersWithBooks(readers map[string]map[string][]string) (count int, err error) {
//
//	var g int
//	for _, v := range readers {
//
//		for _, y := range v {
//			//fmt.Println(x, y, "---")
//			for range y {
//				g++
//
//			}
//		}
//
//	}
//	count = g
//	fmt.Printf("Общее кол-во изданий на руках у каждого читателя - ")
//	return count, nil
//
//}
//
func countBooksEvenReader(readers map[string]map[string][]string) (count int, err error) {

	var g int
	var j int
	for _, v := range readers {
		for _, y := range v {
			if len(y) != 0 {
				j++
				for range y {
					g++
					fmt.Println(g)
				}
			}

		}

	}
	fmt.Println("Количество читателей с изданиями на руках", j)
	count = g
	fmt.Printf("Общее кол-во изданий на руках у каждого читателя - ")
	return count, nil
}
