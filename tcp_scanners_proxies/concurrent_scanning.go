package main

import (
	"fmt"
	"net"
	"sort"
)

// func worker(ports chan int, wg *sync.WaitGroup) {
// 	for p := range ports {
// 		fmt.Println(p)
// 		wg.Done()
// 	}
// }

// func waitGroup() {
// 	var wg sync.WaitGroup
// 	for i := 1; i <= 1024; i++ {
// 		wg.Add(1)
// 		go func(j int) {
// 			defer wg.Done()
// 			address := fmt.Sprintf("scanme.nmap.org:%d", j)
// 			conn, err := net.Dial("tcp", address)
// 			if err != nil {
// 				return
// 			}
// 			conn.Close()
// 			fmt.Printf("%d open\n", j)
// 		}(i)
// 	}
// 	wg.Wait()
// }

// func workerPool() {
// 	ports := make(chan int, 100)
// 	var wg sync.WaitGroup
// 	for i := 0; i < cap(ports); i++ {
// 		go worker(ports, &wg)
// 	}
// 	for i := 1; i <= 1024; i++ {
// 		wg.Add(1)
// 		ports <- i
// 	}
// 	wg.Wait()
// 	close(ports)
// }

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 0; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
