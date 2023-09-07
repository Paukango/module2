package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	in, err := os.Open("data/in.txt")
	if err != nil {
		panic("Open error")
	}
	defer in.Close()

	defer func() {
		if cerr := in.Close(); cerr != nil {
			fmt.Println("Ошибка при закрытии файла:", cerr)
		}
	}()

	sum := fileLen()
	data := make([]byte, sum)
	var j int = 0

	for err != io.EOF {
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
		fmt.Println("Количество строк - ", j)
	}
}

// Подсчет символов в файле
func fileLen() int {
	in, err := os.Open("data/in.txt")
	if err != nil {
		panic("Open error")
	}
	defer in.Close()
	scanner := bufio.NewScanner(in)
	var predLen int
	var sumLen int
	for scanner.Scan() {
		strng := scanner.Text()
		lenStr := len(strng)
		sumLen = sumLen + predLen + lenStr
		predLen = lenStr
	}
	return sumLen
}
