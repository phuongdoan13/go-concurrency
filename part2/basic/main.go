package main

import "sync"

func main() {
  var wg sync.WaitGroup

  wg.Add(3)

  go func() {
    defer wg.Done()
    println("Goroutine 1")
  }()

  go func() {
    defer wg.Done()
    println("Goroutine 2")
  }()

  go func() {
    defer wg.Done()
    println("Goroutine 3")
  }()

  wg.Wait()
}
