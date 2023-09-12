package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	in, err := os.Open("data/in.txt")
	if err != nil {
		log.Print("Open error")
	}

	defer func() {
		if err := in.Close(); err != nil {
			log.Print("Ошибка при закрытии файла:", err)
		}
	}()

	sum, cerr := fileLen() // Обработка ошибок функции подсчета символов в файле
	if cerr != nil && sum == 0 {
		log.Print("Open error")
	} else if cerr != nil && sum > 0 {
		log.Print("Read error")
	}

	data := make([]byte, sum)
	var j int
	for {
		n, err := in.Read(data)
		str := fmt.Sprintf(string(data[:n]))
		strSplt := strings.Split(str, "\n")
		for _, Word := range strSplt {
			if Word != "" {
				j++
			}
		}
		if err == io.EOF {
			break
		}
		fmt.Printf("Total strings: %d", j)
	}
}

// Функция подсчета символов в файле
func fileLen() (int, error) {
	var predLen int
	var sumLen int
	in, err := os.Open("data/in.txt")
	if err != nil {
		log.Print("Open error")
		return sumLen, err
	}

	defer func() {
		if cerr := in.Close(); cerr != nil {
			fmt.Println("Ошибка при закрытии файла:", cerr)
		}
	}()

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		strng := scanner.Text()
		lenStr := len(strng)
		sumLen = sumLen + predLen + lenStr
		predLen = lenStr
		if err != io.EOF && err != nil {
			log.Print("Read error")
			return sumLen, err
		}
	}
	return sumLen, err
}
