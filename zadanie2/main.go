package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("stih.txt")

	data := make([]byte, 64)
	for {
		n, _ := file.Read(data)
		var letters []string
		for i := 0; i < n; i++ {
			letters = append(letters, string(data[i]))
			fmt.Println(letters[i])
		}
	}
}
