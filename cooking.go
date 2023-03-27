package main

import (
	"fmt"
	"sync"
	"time"
)

func main2() {
	started := time.Now()
	foods := []string{"mushroom pizza", "pasta", "kebab", "cake"}
	fmt.Println("started: ", started)
	wg := sync.WaitGroup{}
	wg.Add(len(foods))
	for _, food := range foods {
		go func(food string) {
			cook(food)
			wg.Done()
		}(food)
	}
	wg.Wait()
	fmt.Println("done in", time.Since(started))
}

func cook(food string) {
	fmt.Printf("cooking %s...\n", food)
	time.Sleep(2 * time.Second)
	fmt.Printf("done cooking %s\n", food)
	fmt.Println("")
}
