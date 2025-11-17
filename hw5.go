package main

import "fmt"

func main() {
	fmt.Printf("Задание 1 из ДЗ5: \n")

	numbers := []int{2, 4, 6, 8, 10}
	resultChan := make(chan int, len(numbers))

	for _, num := range numbers {
		go func(n int) {
			square := n * n
			resultChan <- square
		}(num)
	}

	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += <-resultChan
	}

	fmt.Printf("Сумма квадратов чисел: %d", sum)
}
