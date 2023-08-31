package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {

	defer executionTime()()
	inFile, err := os.Open("06_task_in.txt")
	if err != nil {
		log.Fatal("Не найден файл")
	}
	defer inFile.Close()
	outFile, er := os.Create("Out.txt")
	if er != nil {
		log.Fatal("Ошибка создания или открытия файла")
	}
	defer outFile.Close()
	i := 1
	s := bufio.NewScanner(inFile)
	for s.Scan() {
		number := strconv.Itoa(i)
		outFile.WriteString(number)
		outFile.WriteString("  ")
		outFile.WriteString(s.Text())
		outFile.WriteString("\n")
		i++
	}
}
func executionTime() func() {
	startTime := time.Now()
	return func() {
		elapsedTime := time.Since(startTime)
		fmt.Printf("Функция выполнилась за %v\n", elapsedTime)
	}
}
