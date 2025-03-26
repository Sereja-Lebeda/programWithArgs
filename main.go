package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Получаем аргументы командной строки (первый - это сам исполняемый файл)
	args := os.Args

	// Если аргументов нет, выводим сообщение
	if len(args) < 2 {
		fmt.Println("Использование: program <параметры>")
		return
	}

	// Выводим все переданные аргументы
	fmt.Println("Переданные параметры:")
	for i, arg := range args[1:] { // Пропускаем args[0] (имя исполняемого файла)
		fmt.Printf("Аргумент %d: %s\n", i+1, arg)
	}
	time.Sleep(30 * time.Second)
}
