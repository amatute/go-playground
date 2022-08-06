package main

import (
	"fmt"
	"sync"
	"time"
)

type Store struct {
	id      string
	name    string
	number  string
	ranking int
}

var stores = []Store{
	{name: "store A", number: "1", ranking: 1},
	{name: "store B", number: "2", ranking: 2},
	{name: "store C", number: "3", ranking: 3},
	{name: "store D", number: "4", ranking: 4},
	{name: "store E", number: "5", ranking: 5},
	{name: "store F", number: "6", ranking: 6},
	{name: "store G", number: "7", ranking: 7},
	{name: "store H", number: "8", ranking: 8},
}

func main() {
	fmt.Println("GO patterns!")
	c := gen(stores...)

	c1 := rankStore(c)
	c2 := rankStore(c)
	c3 := rankStore(c)

	out := fanIn(c1, c2, c3)

	for processed := range out {
		fmt.Println("Name:", processed.name, "Number:", processed.number, "Ranking:", processed.ranking)
	}
}

func gen(stores ...Store) <-chan Store {
	out := make(chan Store, len(stores))
	defer close(out)
	for _, s := range stores {
		out <- s
	}
	return out
}

func rankStore(stores <-chan Store) <-chan Store {
	out := make(chan Store)
	go func() {
		defer close(out)
		for i := range stores {
			time.Sleep(time.Second / 2) // block
			if i.ranking%2 == 0 {
				i.ranking += 100
			}
			out <- i
		}
	}()
	return out
}

func fanIn(channels ...<-chan Store) <-chan Store {
	var wg sync.WaitGroup
	out := make(chan Store)
	output := func(c <-chan Store) {
		defer wg.Done()
		for i := range c {
			out <- i
		}
	}
	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}
	go func ()  {
			wg.Wait()
			close(out)
	}()
	return out
}
