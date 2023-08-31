package main

import "fmt"

func main() {

	slice := []string{"Один", "Два", "Три"}
	s := []int{0, 2, 0, 0, 1}
	var (
		word string
		ok   bool
	)
	ok = contains(slice, word)
	fmt.Println(ok)
	//--------------------
	maxArg := getMax(s...)
	fmt.Println("Максимальное число = ", maxArg)
}

func contains(slice []string, word string) bool {
	fmt.Println("Введите строку:")
	fmt.Scan(&word)
	for i := 0; i < len(slice); i++ {
		if slice[i] == word {
			return true
		}
	}
	return false

}
func getMax(args ...int) int {
	var d int
	var j int
	for i := range args {
		x := args[i]
		if x >= d {
			j = x
		} else if x < d {
			continue
		}
		d = x
	}
	return j
}
