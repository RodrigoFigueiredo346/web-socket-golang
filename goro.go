package main

import (
	"fmt"
	"sync"
	"time"
)

func Goro() {

	start := time.Now()

	ch := make(chan string, 3)

	var wg sync.WaitGroup

	wg.Add(2)

	go user(&wg, ch)
	go card(&wg, ch)
	go limit(ch)

	//sem wg os canais informam em tempo real, com wg espera todos comunicarem e executa e resto do código
	wg.Wait()

	//sem close nunca passa daqui, com close nunca executa o limit
	//close(ch)

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println(time.Since(start))
}

func user(wg *sync.WaitGroup, ch chan string) {
	time.Sleep(time.Millisecond * 1000)
	ch <- "Rodrigo"
	wg.Done()
}
func card(wg *sync.WaitGroup, ch chan string) {
	time.Sleep(time.Millisecond * 2000)
	ch <- "Visa"
	wg.Done()
}

func limit(ch chan string) {

	for {
		time.Sleep(time.Millisecond * 3000)
		ch <- "2000"
	}
}
