package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SyncedData struct {
	inner map[string]int
	mutex sync.Mutex
}

func (d *SyncedData) Insert(k string, v int) {
	d.mutex.Lock()
	d.inner[k] = v
	d.mutex.Unlock()
}

func (d *SyncedData) Get(k string) int {
	d.mutex.Lock()
	data := d.inner[k]
	d.mutex.Unlock()
	return data
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wait sync.WaitGroup
	counter := 0

	for i := 0; i < 5; i++ {
		wait.Add(1)

		counter += 1
		go func() {
			defer func() {
				fmt.Println(counter, "goroutine remaining")
				counter -= 1
				wait.Done()
			}()
			duration := time.Duration(rand.Intn(500) * int(time.Millisecond))
			fmt.Println("waiting for", duration)
			time.Sleep(duration)
		}()
	}
	fmt.Println("waiting for goroutines to finish")
	wait.Wait()
	fmt.Println("done!")
	// Data := SyncedData{
	// 	inner: map[string]int{},
	// }
	// Data.Insert("sample", 2)
	// Data.Insert("test", 5)
	// Data.Get("sample")
	// Data.Get("test")
}
