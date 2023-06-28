package main

import (
  "time"
  "fmt"
  "sync"
)

func main() {
  now := time.Now()
  var wg sync.WaitGroup
  wg.Add(1)
  go func() {
    defer wg.Done()
    work()
  }() 

  wg.Wait()
  fmt.Println("elapsed", time.Since(now))
  fmt.Println("Done")
}

func work() {
  time.Sleep(500 * time.Millisecond)
  fmt.Println("Work done")
}
