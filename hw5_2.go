package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Нужно ввести: go run main.go <количество_воркеров>")
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		log.Fatal("Количество воркеров должно быть положительным числом")
	}

	fmt.Printf("Запуск %d воркеров\n", numWorkers)
	ctx, cancel := context.WithCancel(context.Background())
	dataChan := make(chan string, 100)
	var wg sync.WaitGroup
	var once sync.Once
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, dataChan, &wg)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer once.Do(func() { close(dataChan) })
		produceData(ctx, dataChan)
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		fmt.Println("\nПолучен сигнал завершения...")
		cancel()
	}()
	wg.Wait()
	fmt.Println("Все горутины завершены. Выход.")
}

func worker(ctx context.Context, id int, dataChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d: завершение по контексту\n", id)
			return
		case data, ok := <-dataChan:
			if !ok {
				fmt.Printf("Воркер %d: канал закрыт\n", id)
				return
			}
			fmt.Printf("Воркер %d получил: %s\n", id, data)
		}
	}
}

func produceData(ctx context.Context, dataChan chan<- string) {
	counter := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Продюсер: остановка генерации данных")
			return
		default:
			data := fmt.Sprintf("сообщение-%d", counter)
			select {
			case dataChan <- data:
				counter++
			case <-ctx.Done():
				return
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
}
