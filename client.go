package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"sort"
	"sync"
	"time"
)

var (
	Num  int
	Conn int
	Addr string
)

const RS = 100000

func init() {
	flag.IntVar(&Num, "num", 1, "")
	flag.IntVar(&Conn, "conn", 1, "")
	flag.StringVar(&Addr, "addr", "192.168.101.232", "")
	flag.Parse()
}

func main() {
	clients := make([]*rpc.Client, 0, 10)
	for i := 0; i < Conn; i++ {
		client, err := rpc.Dial("tcp", Addr+":8000")
		if err != nil {
			log.Fatalln(err)
		}
		clients = append(clients, client)
	}
	index := 0

	s1 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	var wg sync.WaitGroup

	start := time.Now()

	d := make([]float64, 0)
	mx := sync.Mutex{}
	for i := 0; i < Num; i++ {
		wg.Add(1)

		var c *rpc.Client
		if index < len(clients) {
			c = clients[index]
			index++
		} else {
			index = 0
			c = clients[index]
		}

		go func(cli *rpc.Client) {
			defer wg.Done()

			log.Println("goroutine start...")
			st := time.Now()

			for n := 0; n < RS; n++ {
				t1 := time.Now().UnixNano()
				if err := cli.Call("EchoServer.Echo", &s1, nil); err != nil {
					log.Println(err)
				}
				t2 := time.Now().UnixNano()
				mx.Lock()
				d = append(d, float64(t2 - t1) / 1e6)
				mx.Unlock()
			}

			log.Println(time.Now().Sub(st))
		}(c)
	}
	wg.Wait()

	total := RS * Num
	secs := float64(time.Now().Sub(start)) / 1000000000

	fmt.Printf("concurrency: %d\n", Num)
	fmt.Printf("total: %d\n", total)
	fmt.Printf("seconds: %v\n", secs)
	fmt.Printf("qps: %d\n", int64(float64(total)/(secs)))

	sort.Float64s(d)
	totalTime := 0.0
	for _, t := range d{
		totalTime += t
	}
	fmt.Printf("avg: %v\n", totalTime / float64(total))
	fmt.Printf("50%%: %v\n", d[total * 50/100])
	fmt.Printf("99%%: %v\n", d[total * 99/100])
	fmt.Printf("100%%: %v\n", d[total - 1])
}