package main

import (
	"fmt"
	"time"
)

/*
 Links from YouTube tutorials:
 Maria Carrion: https://www.youtube.com/watch?v=x6vBvgKGvxU
 Katan Coding:  https://www.youtube.com/watch?v=qyM8Pi1KiiM
 Golang Cafe:   https://www.youtube.com/watch?v=rlPbtzBmHp8&list=PLlhUYfyFOzo2QTD5z4LUwd13-_9Pwh028
 https://www.youtube.com/watch?v=245QljYu3-A
*/
type Store struct {
	name    string
	number  string
	ranking int
}

type Report struct {
	storeNbr       string
	productNumber  string
	priceEcommerce string
	priceLocal     string
}

func main() {
	defer timeTrack(time.Now(), "process")
	fmt.Println("GO concurrecy patterns!")

	stores := GetStores(100)

	//fmt.Println(stores)
	c := gen(stores...)

	c1 := DoReport(c)
	c2 := DoReport(c)
	c3 := DoReport(c)
	c4 := DoReport(c)

	//c1 := rankStore(done, c)
	//c2 := rankStore(done, c)
	//c3 := rankStore(done, c)

	out := fanInMerge(c1,c2,c3,c4)

	for c := range out {
		fmt.Println(c)
	}
	
	// for {
	// 	if c1 == nil && c2 == nil {
	// 		break
	// 	}
		
	// 	select {
	// 	case _, ok := <- c1:
	// 		if !ok {
	// 			c1 = nil
	// 		}
	// 	case _, ok := <- c2:
	// 		if !ok {
	// 			c2 = nil
	// 		}
	// 	}
	// }
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", name, elapsed)
}

func gen(stores ...[]Store) <-chan []Store {
	out := make(chan []Store, len(stores))

	go func() {
		for _, s := range stores {
			out <- s
		}
		close(out)
	}()
	return out
}

func rankStore(done <-chan bool, stores <-chan Store) <-chan Store {
	out := make(chan Store)
	go func() {
		defer close(out)
		for i := range stores {
			time.Sleep(time.Second / 2) // block
			if i.ranking%2 == 0 {
				i.ranking += 100
			}
			select {
			case out <- i:
			case <-done:
				return
			}
		}
	}()
	return out
}

func DoReport(stores <-chan []Store) <-chan Report {
	out := make(chan Report)

	go func() {
		for chunk := range stores {
			// do heavy lifting
			for _, s := range chunk {
				time.Sleep(time.Millisecond * 250)
				out <- Report{storeNbr: s.number, productNumber: "PROD_100", priceEcommerce: "5990", priceLocal: "9990"}
			}
		}
		close(out)
	}()

	return out
}

func fanInMerge(cs ...<-chan Report) <-chan Report {
	chans := len(cs)
	wait := make(chan struct{}, chans)

	out := make(chan Report)

	send := func(reports <-chan Report) {
		defer func() { wait <- struct{}{} }()

		for r := range reports {
			out <- r
		}
	}

	for _, c := range cs {
		go send(c)
	}

	go func() {
		for range wait {
			fmt.Println(chans)
			chans--
			if chans == 0 {
				break
			}
		}
		close(out)
	}()

	return out
}


// Fan in with Wait Grup
// func fanIn(done <-chan bool, channels ...<-chan Store) <-chan Store {
// 	var wg sync.WaitGroup
// 	out := make(chan Store)

// 	output := func(c <-chan Store) {
// 		defer wg.Done()
// 		for i := range c {
// 			select {
// 			case out <- i:
// 			case <-done:
// 				return
// 			}
// 		}
// 	}

// 	wg.Add(len(channels))
// 	for _, c := range channels {
// 		go output(c)
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()
// 	return out
// }

func GetStores(n int) [][]Store {
	stores := []Store{}
	for i := 0; i < n; i++ {
		s := Store{
			name:    "Store",
			number:  fmt.Sprintf("%d", i),
			ranking: i,
		}
		stores = append(stores, s)
	}
	return chunkSlice(stores, 3)
}

func chunkSlice(slice []Store, chunkSize int) [][]Store {
	var chunks [][]Store
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}
