package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
)

type Task struct {
	ID        int    `json:"id"`
	UserId    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main3() {
	var t Task
	wg := &sync.WaitGroup{}
	wg.Add(100)
	// create a semaphore with a buffer of 10, only 10 go routines can run at a time
	// this is a buffered channel
	// the buffer is the number of go routines that can run at a time
	// will block the rest of the code until a place is available in the semaphore if we try to exceed the buffer of 10
	sem := make(chan bool, 10)

	for i := 0; i < 100; i++ {
		fmt.Println(runtime.NumGoroutine())
		sem <- true

		go func(i int) {
			defer wg.Done()
			defer func() { <-sem }() // reserve a place in the semaphore, will block rest of code until a place is available in the semaphore
			res, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", i))
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
				log.Fatal(err)
			}
			fmt.Println(t.Title)
		}(i)
	}
	wg.Wait()
}
