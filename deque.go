package main

import "fmt"

type Deque interface {
	PushBack(int)
	PushFront(int)
	Size() int
	Front() int
	Back() int
	PopFront() int
	PopBack() int
}

type DequeRealisation struct {
	DequeSample []int
}

func NewDeque() *DequeRealisation {
	return &DequeRealisation{
		DequeSample: make([]int, 0),
	}
}

func (object *DequeRealisation) PushBack(elem int) {
	object.DequeSample = append(object.DequeSample, elem)
}

func (object *DequeRealisation) PushFront(elem int) {
	object.DequeSample = append([]int{elem}, object.DequeSample...)
}

func (object *DequeRealisation) Size() int {
	return len(object.DequeSample)
}

func (object *DequeRealisation) Front() int {
	if object.IsEmpty() {
		panic("Очередь пустая!")
	}
	return object.DequeSample[0]
}

func (object *DequeRealisation) Back() int {
	if object.IsEmpty() {
		panic("Очередь пустая!")
	}
	return object.DequeSample[len(object.DequeSample)-1]
}

func (object *DequeRealisation) String() string {
	return fmt.Sprintf("%v", object.DequeSample)
}

func (object *DequeRealisation) IsEmpty() bool {
	if len(object.DequeSample) == 0 {
		return true
	} else {
		return false
	}
}

func (object *DequeRealisation) PopFront() int {
	if object.IsEmpty() {
		panic("Очередь пустая!")
	}
	front := object.Front()
	object.DequeSample = object.DequeSample[1:]
	return front
}

func (object *DequeRealisation) PopBack() int {
	if object.IsEmpty() {
		panic("Очередь пустая!")
	}
	backIndex := len(object.DequeSample) - 1
	back := object.Back()
	object.DequeSample = object.DequeSample[:backIndex]
	return back
}

func (object *DequeRealisation) Clear() {
	object.DequeSample = make([]int, 0)
}

func main() {
	deque := NewDeque()
	fmt.Print("Создали пустую очередь при помощи конструктора\n\n")
	fmt.Printf("Проверка на пустоту: %v\n", deque.IsEmpty())
	fmt.Print("Добавили элементы в конец: 1, 2, 3\n")
	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushBack(3)

	fmt.Printf("Очередь: %v\n\n", deque)
	fmt.Printf("Размер: %d\n", deque.Size())
	fmt.Printf("Первый элемент: %d\n\n", deque.Front())
	popped := deque.PopFront()
	fmt.Printf("Удаленный элемент из начала: %d\n", popped)
	fmt.Printf("Очередь после удаления элемента: %v\n\n", deque)

	fmt.Print("Добавили элементы в начало: 4, 5, 6\n")
	deque.PushFront(4)
	deque.PushFront(5)
	deque.PushFront(6)
	fmt.Printf("Очередь: %v\n\n", deque)

	fmt.Print("Удаляем элемент с конца\n")
	popped = deque.PopBack()
	fmt.Printf("Удаленный элемент: %d\n\n", popped)
	fmt.Printf("Очередь после удаления элемента: %v\n\n", deque)

	fmt.Print("Очищаем очередь\n")
	deque.Clear()
	fmt.Printf("Проверка на пустоту после очищения: %v\n", deque.IsEmpty())
	fmt.Printf("Размер очищенной очереди: %d\n", deque.Size())
}
