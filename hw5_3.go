package main

import (
	"fmt"
	"sync"
)

type MyMap struct {
	mutex  sync.RWMutex
	mapObj map[string]int
}

func NewMap() *MyMap {
	return &MyMap{
		mapObj: make(map[string]int),
	}
}

func (myMap *MyMap) Set(key string, value int) {
	myMap.mutex.Lock()
	defer myMap.mutex.Unlock()
	myMap.mapObj[key] = value
}

func (myMap *MyMap) Get(key string) (int, bool) {
	myMap.mutex.RLock()
	defer myMap.mutex.RUnlock()
	val, ok := myMap.mapObj[key]
	return val, ok
}

func (myMap *MyMap) Delete(key string) {
	myMap.mutex.Lock()
	defer myMap.mutex.Unlock()
	delete(myMap.mapObj, key)
}

func main() {
	myMap := NewMap()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			myMap.Set(fmt.Sprintf("key%d", id), id*10)
		}(i)
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		if val, ok := myMap.Get(fmt.Sprintf("key%d", i)); ok {
			fmt.Printf("key%d = %d\n", i, val)
		}
	}
}
