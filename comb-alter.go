package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Программа объединяет два списка, чередуя элементы")

	//Запрашиваем первый список и записываем в переменную
	fmt.Println("Введите первый список через пробел")
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	first := myscanner.Text()

	//Запрашиваем второй список и записываем в переменную
	fmt.Println("Введите второй список через пробел")
	myscanner1 := bufio.NewScanner(os.Stdin)
	myscanner1.Scan()
	second := myscanner1.Text()

	//Делим строки на массивы для дальнейшего объединения
	var firstSplit = strings.Split(first, " ")
	var secondSplit = strings.Split(second, " ")

	if len(firstSplit) == len(secondSplit) {
		for i := 0; i < len(firstSplit); i++ {
			firstSplit[i] = firstSplit[i] + " " + secondSplit[i]
		}
	} else {
		fmt.Println("Неравное количество элементов в списках, введите заново")
	}
	fmt.Println(firstSplit)
}
