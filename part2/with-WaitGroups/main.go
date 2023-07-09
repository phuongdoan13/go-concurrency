package main

import (
  "fmt"
  "runtime"
  "time"
	"sync"
)

func main() {
  fmt.Println("number of cores", runtime.NumCPU())
  var wg sync.WaitGroup
	wg.Add(10)
	now := time.Now()
  for i:=0; i<10; i++ {
    go work(&wg, i+1)
  }
  wg.Wait()
	fmt.Println(time.Since(now))
  fmt.Println("main function")
}

func work(wg *sync.WaitGroup, id int) {
  defer wg.Done()
  time.Sleep(100 * time.Millisecond)
  fmt.Println("work", id, "is done")
}
