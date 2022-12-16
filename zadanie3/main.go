package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	var surname string
	fmt.Print("Введите имя: ")
	fmt.Fscan(os.Stdin, &name)

	fmt.Print("Введите фамилию: ")
	fmt.Fscan(os.Stdin, &surname)

	fmt.Println(name, surname)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите заметку: ")
	zametka, _ := reader.ReadString('\n')
	fmt.Println(name, surname)
	fmt.Println(zametka)

