package main

import "fmt"

const maxNum = 12306

func main() {

	fmt.Printf("Введите число, не превышающее %d:  ", maxNum)
	var number int
	fmt.Scan(&number)

	if number > maxNum {
		fmt.Println("Ошибка. Введённое число превышает границу")

	} else {
		var result int = execute(number)
		if result != 0 {
			fmt.Printf("Ответ: %d", result)
		}
	}

}

func execute(num int) int {
	var iteration int = 0
	for num <= maxNum {
		if iteration > 1000 {
			fmt.Println("Слишком много итераций. Текущее число: ", num)
			break
		}

		if num < 0 {
			num *= -1
		} else if num%7 == 0 {
			num *= 39
		} else if num%9 == 0 {
			num = num*13 + 1
			iteration++
			continue
		} else {
			num = (num + 2) * 3
		}

		if num%13 == 0 && num%9 == 0 {
			fmt.Printf("service error. \nТекущее число: %d", num)
			return 0
		} else {
			num += 1
		}
		iteration++
	}
	return num
}
