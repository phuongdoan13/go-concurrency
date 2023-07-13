package main

import (
	"sync"
	"net"
	"io/ioutil"
	"log"
	"fmt"
)

func main() {
	total, max := 10, 3
	var wg sync.WaitGroup

	for i := 0; i < total; i+=max{
		limit := max
		if i+max > total {
			limit = total - i
		}

		wg.Add(limit)
		for j := 0; j < limit; j++ {
			go func(j int) {
				defer wg.Done()
				conn, err := net.Dial("tcp", ":8080")
				if err != nil {
					log.Fatal(err)
				}

				bs, err := ioutil.ReadAll(conn)
				if err != nil {
					log.Fatal(err)
				}

				if string(bs) != "success" {
					log.Fatal("failed")
				}

				fmt.Printf("request %d: success\n", i + 1 + j)
			}(j)
		}
		wg.Wait()
	}
}
