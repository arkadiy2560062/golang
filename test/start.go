package main

import "fmt"

func main() {

	fmt.Println("Привет, Мир! Привет, Кодбай!")

	names := []string{"я", "балдею", "от", "Го"} // срез строк

	for _, i := range names { //проходим в цикле печатая каждое слово
		fmt.Println(i)

	}

}
