package main

import (
	"fmt"
	"log"

	"os"
)

func main() {
	a := make([]byte, 0, 0)
	ReadFile(a)
	fmt.Println(string(a))

}

func ReadFile([]byte) []byte {

	fileData, er := os.ReadFile("06_task_in")
	if er != nil {
		log.Fatal()
		fmt.Println("File not opend")
	}
	fmt.Println(string(fileData))
	return fileData
}

func WriteFile() {

	data := []byte("Текст")
	err := os.WriteFile("Tex.txt", data, 0600)
	if err != nil {
		log.Fatal()
		fmt.Println("File not create")
	}
}
