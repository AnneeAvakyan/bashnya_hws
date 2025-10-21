package main

import (
	"fmt"
)

const maxNum = 12306

var from1To19 = []string{"", "один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять", "десять",
	"одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать"}
var less100 = []string{"десять", "двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто"}
var thousands = []string{"", "тысяча", "тысячи", "тысяч"}

func main() {
	fmt.Printf("Введите число, не превышающее %d:  ", maxNum)
	var number int
	fmt.Scan(&number)
	if number > maxNum {
		fmt.Println("Ошибка. Введённое число превышает границу")

	} else {
		var result int = execute(number)
		if result != 0 {
			fmt.Printf("Ответ: %d\n", result)
			fmt.Println(makeWords(result))
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

func makeWords(num int) string {
	result := ""
	num, result = convertEnsOfThous(num, result)
	num, result = convertThous(num, result)
	num, result = convertHundreds(num, result)
	result = convertDoubleDigits(num, result)
	return result
}

func convertEnsOfThous(num int, result string) (int, string) {
	if num >= 20000 && num < 100000 {
		ending := (num / 1000) % 10
		if ending == 0 {
			result += less100[(num/10000)-1] + " " + thousands[3] + " "
		} else if ending == 1 {
			result += less100[(num/10000)-1] + " "
		} else if ending == 2 {
			result += less100[(num/10000)-1] + " "
		} else {
			result += less100[(num/10000)-1] + from1To19[(num/100000)%10] + " "
		}
		num %= 10000
	}
	return num, result
}

func convertThous(num int, result string) (int, string) {
	if num >= 1000 && num < 20000 {
		ending := (num / 1000) % 10
		if ending == 1 && (num/10000)%10 != 1 {
			result += "одна" + " " + thousands[1] + " "
		} else if ending >= 2 && ending <= 4 && (num/10000)%10 != 1 {
			result += from1To19[(num/1000)] + " " + thousands[2] + " "
		} else {
			result += from1To19[(num/1000)] + " " + thousands[3] + " "
		}
		num %= 1000
	}
	return num, result
}

func convertHundreds(num int, result string) (int, string) {
	if num >= 100 {
		ending := (num / 100) % 10
		switch ending {
		case 1:
			result += "сто" + " "
		case 2:
			result += "двести" + " "
		case 3:
			result += "триста" + " "
		case 4:
			result += "четыреста" + " "
		case 5, 6, 7, 8, 9:
			result += from1To19[num/100] + "сот" + " "
		}
		num %= 100
	}
	return num, result
}

func convertDoubleDigits(num int, result string) string {
	if num >= 20 {
		result += less100[(num/10)-1] + " "
		num %= 10
	}
	result += from1To19[num]
	return result
}
