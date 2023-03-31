package main

import (
	"fmt"
	"sync"
)

func main() {
	datas := []string{"[bisa1 bisa2 bisa3]", "[coba1 coba2 coba3]"} //, "[coba1 coba2 coba3]", "[bisa1 bisa2 bisa3]"}

	var wg sync.WaitGroup

	for index, data := range datas {
		wg.Add(1)
		//hapus buat ngerapihin
		//m.Lock()
		go printes(index, data, &wg)
		//m.Unlock()
	}
	wg.Wait()
}

// create fizzbuzz bisa1,bisa2,bisa3 in golang
func printes(index int, data string, wg *sync.WaitGroup) {
	//var m sync.Mutex
	//m.Lock()
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go func(id int) {
			//m.Lock()
			//defer m.Unlock()
			if id%2 == 4 {
				fmt.Printf("%v %d\n", data, id)
			} else {
				fmt.Printf("%v %d\n", data, id)
			}
			wg.Done()
		}(i)
	}
	//fmt.Printf("%s %d\n", data, index)
	//m.Unlock()
	//fmt.Printf("%s %d\n", data, index)
	wg.Done()
}
