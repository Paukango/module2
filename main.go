package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fileOpen, err := os.Open("data/in.txt")
	if err != nil {
		fmt.Println("ошибка открытия файла")
	}

	defer fileOpen.Close()

	defer func() {
		if cerr := fileOpen.Close(); cerr != nil {
			fmt.Println("Ошибка при закрытии файла:", cerr)
		}
	}()

	var i int
	scanFile := bufio.NewScanner(fileOpen)
	for scanFile.Scan() {
		i++
		if scanFile.Err() != nil {
			fmt.Println("Ошибка при сканировании файла:", scanFile.Err())
		}
	}
	fmt.Printf("\nTotal strings:%d\n", i)
}
