package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Открываем файл для чтения
	file, err := os.Open("data/07_task_in.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	fileOut, err := os.Create("data/data_out.txt")
	if err != nil {
		fmt.Println("Ошибка при создании или открытии файла:", err)
		return
	}
	defer file.Close()

	defer func() {
		fileOut, err := os.Open("data/data_out.txt")
		if err != nil {
			fmt.Println("Ошибка при создании или открытии файла:", err)
			return
		}
		scan := bufio.NewScanner(fileOut)
		for scan.Scan() {
			scanFile := scan.Text()
			fmt.Println(scanFile)
		}
	}()

	// Создаем сканер для чтения из файла
	scanner := bufio.NewScanner(file)
	var stringNum int

	for scanner.Scan() {
		workSt := scanner.Text()
		workSlice := strings.SplitN(workSt, "|", 3)

		stringNum++

		name := workSlice[0]
		address := workSlice[1]
		city := workSlice[2]

		stringOut := fmt.Sprintf("Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n", stringNum, name, address, city)

		fileOut.WriteString(stringOut)

		for _, word := range workSlice {
			if word == "" {
				err := fmt.Sprintf("parse error: empty field on string %d", stringNum)
				panic(err)
			}

		}

	}

}
