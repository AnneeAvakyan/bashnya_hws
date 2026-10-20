package main

import "fmt"

// multiply — умножает каждое значение на 3 и возвращает канал с результатами
func multiply(doneCh chan struct{}, inputCh chan int) chan int {
	resultCh := make(chan int)

	go func() {
		defer close(resultCh)

		for value := range inputCh {
			result := value * 2

			select {
			case <-doneCh:
				return
			case resultCh <- result:
			}
		}
	}()

	return resultCh
}

// generator — отправляет данные в канал
func generator(doneCh chan struct{}, numbers []int) chan int {
	outputCh := make(chan int)

	go func() {
		defer close(outputCh)

		for _, num := range numbers {
			select {
			case <-doneCh:
				return
			case outputCh <- num:
			}
		}
	}()

	return outputCh
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Исходный массив: (данные для канала inputCh) %v\n", numbers)

	doneCh := make(chan struct{})
	defer close(doneCh)

	fmt.Println("Запускаем генератор, который отправляет числа")
	inputCh := generator(doneCh, numbers)
	fmt.Println("Записываем в другой канал результат умнножения на 2")
	resultCh := multiply(doneCh, inputCh)

	for res := range resultCh {
		fmt.Println(res)
	}
}
