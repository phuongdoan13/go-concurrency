package main

import (
  "time"
  "fmt"
)

func main() {
  now := time.Now()
  done := make(chan struct{})
  go task1(done)
  go task2(done)
  go task3(done)
  go task4(done)
  <- done
  fmt.Println("elapsed", time.Since(now))
  fmt.Println("Done")
}

func task1(done chan struct{}) {
  time.Sleep(100 * time.Millisecond)
  fmt.Println("task1")
  done <- struct{}{}
}

func task2(done chan struct{}) {
  time.Sleep(100 * time.Millisecond)
  fmt.Println("task2")
  done <- struct{}{}
}

func task3(done chan struct{}) {
  time.Sleep(100 * time.Millisecond)
  fmt.Println("task3")
  done <- struct{}{}
}

func task4(done chan struct{}) {
  time.Sleep(100 * time.Millisecond)
  fmt.Println("task4")
  done <- struct{}{}
}
