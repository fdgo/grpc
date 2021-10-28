package main

import (
	"net/http"
	"sync"
)

func main() {
	//X()
	var wg sync.WaitGroup
	wg.Add(20000)
	for i := 0; i < 20000; i++ {
		go TestSell(&wg)
	}
	wg.Wait()
}
func TestSell(wg *sync.WaitGroup) {
	defer wg.Done()
	http.Get("http://192.168.164.128:8080/ping")
}
func X()  {
	http.Get("http://192.168.164.128:8080/ping")
}