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
	outFile, err := os.Create("data/out.txt")
	if err != nil {
		fmt.Println("ошибка создания файла")
	}

	scanFile := bufio.NewScanner(fileOpen)
	for scanFile.Scan() {
		reader := scanFile.Text()
		fmt.Println(reader)
		outFile.WriteString(scanFile.Text())
		outFile.WriteString("\n")

	}

}
