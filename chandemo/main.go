package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan string, 10)
	go consumer(ch, wg)
	//go producer(ch, wg)
	wg.Wait()
}

func consumer(ch chan string, wg *sync.WaitGroup) {
	index := false
	for {
		if !index {
			close(ch)
			index = true
		}
		fmt.Println("aaaa")
		str := <-ch
		fmt.Println(str)
		fmt.Println("aaaa2")
	}
	fmt.Println("aaaa3")
	wg.Done()
}

func producer(ch chan string, wg *sync.WaitGroup) {
	index := 0
	for {
		index = index + 1
		ch <- strconv.Itoa(index)
		time.Sleep(time.Second)
		if index > 10 {
			close(ch)
			return
		}
	}
	wg.Done()
}
