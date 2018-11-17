package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, c chan int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Println("Worker ID:", id, ",received ", string(n))
		go func() { wg.Done() }()
	}
}

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWorker(id, w.in, wg)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup
	/*添加要等待的任务数*/
	wg.Add(20)

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
		//wg.Add(1)
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
		//<-workers[i].done
	}
	wg.Wait()
}

func main() {
	chanDemo()
}
