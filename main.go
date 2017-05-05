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
var m sync.Mutex

func main() {
	rand.Seed(time.Now().UnixNano())

	l := 100
	a := l

	tasks := make(chan Task, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < l; i++ {
			Println(i)
			select {
			case tasks <- Task{"1", rand.Float64()}:
			default:
				i--
				// Println("OVERFLOW, WAITING")
				time.Sleep(300 * time.Millisecond)
			}
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
					m.Lock()
					a--
					m.Unlock()
					if task.Attemts >= 0.5 {
						task.Attemts -= .1
						tasks <- task
						m.Lock()
						a++
						m.Unlock()
					}
				}
			default:
				time.Sleep(time.Millisecond)
				if a == 0 {
					return
				}
				continue
			}
		}
	}()

	wg.Wait()
	Println(a)
}
