package main

import (
	. "fmt"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	Url     string
	Attemts float64
}

var wg sync.WaitGroup

func main() {

	tasks := make(chan Task, 100)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			tasks <- Task{"1", rand.Float64()}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case task, ok := <-tasks:
				if ok {
					Println(task)
					if task.Attemts >= 0.5 {
						task.Attemts -= .1
						tasks <- task
					}
				}
			default:
				time.Sleep(time.Millisecond)
				continue
			}
		}
	}()

	wg.Wait()
}
