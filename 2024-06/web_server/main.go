package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

var data [][]byte

func updateData(wg *sync.WaitGroup) {
	defer wg.Done()
	for range 10 {
		data = append(data, make([]byte, 1024*1024))
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {

	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var wg sync.WaitGroup
		wg.Add(1)
		go updateData(&wg)
		wg.Wait()
		w.Write([]byte("ok"))
	})

	// 开启http服务
	log.Println("server start at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
