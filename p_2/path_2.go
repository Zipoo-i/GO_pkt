package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Printf("ERRORS: Ошибка открытия\n%v\n", err)
		return
	}
	defer func(){
		if err := file.Close(); err != nil{
			fmt.Printf("ERRORS: Ошибка закрытия\n%v\n", err)
		}
	}()
	reading(file)
}

func reading(file *os.File) {
	read := bufio.NewReader(file)
	for{
		lines, err := read.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Printf("ERRORS: Ошибка чтения\n%v\n", err)
			}
			break
		}
		PrintH(lines)
	}
}

func PrintH(lines string){
	fmt.Print(lines)
}