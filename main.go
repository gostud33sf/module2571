package main

import (
	"fmt"
	"os"
)

func main() {
	var n string
	fmt.Print("Введите данные: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
		os.Exit(1)
	}
	fmt.Printf("Вы ввели следующие данные: %s\n", n)
}
