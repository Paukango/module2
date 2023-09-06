package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	fileOpen, err := os.Open("data/in.txt")
	if err != nil {
		fmt.Println("ошибка открытия файла")
	}
	defer fileOpen.Close().Error()

	var i int
	buffer := make([]byte, 20)
	scanFile := bufio.NewScanner(fileOpen)
	reader := strings.NewReader(scanFile.Text())

	for scanFile.Scan() {
		i++
		_, err := reader.Read(buffer)
		if err == io.EOF {

			break
		}
		if err != nil {
			fmt.Printf("Ошибка чтения: %v\n", err)
			break
		}

		fmt.Println(reader)

	}
	fmt.Printf("Total strings:%d", i)
}
