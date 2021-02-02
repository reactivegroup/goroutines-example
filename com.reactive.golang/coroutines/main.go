package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// goroutines count
	// Friendly hint: too much will lead to OOM
	var goroutinesCount = 100

	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < goroutinesCount; i++ {
		go func(i int) {
			wg.Add(1)
			for j := 0; j < 10; j++ {
				// open debug log
				// fmt.Println("Go-" + strconv.Itoa(i) + " loop-" + strconv.Itoa(j))
				time.Sleep(time.Duration(1) * time.Microsecond)
			}
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("App elapsed: ", elapsed)
}
