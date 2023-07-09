package main

import (
  "fmt"
  "runtime"
  "time"
)

func main() {
  // If the Sleep is too short, not all the work will be done
  // Need to use WaitGroup to wait for all the work to be done
  fmt.Println("number of cores", runtime.NumCPU())
  for i:=0; i<10; i++ {
    go work(i+1)
  }
  time.Sleep(100 * time.Millisecond)
  fmt.Println("main function")
}

func work(id int) {
  time.Sleep(100 * time.Millisecond)
  fmt.Println("work", id, "is done")
}
