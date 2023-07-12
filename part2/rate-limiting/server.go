package main

import (
	"log"
	"net"
	"sync/atomic"
	"time"
)

func main(){
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Could not create listener: %v", err)
	}

	var connections int32
	for {
		conn, err := li.Accept()
		if err != nil {
			continue
		}
		connections++

		go func(con net.Conn) {
			defer func() {
				_ = con.Close()
				atomic.AddInt32(&connections, -1)
			}()

			if atomic.LoadInt32(&connections) > 3 {
				return
			}

			time.Sleep(time.Second)
			_, err := con.Write([]byte("success"))
			if err != nil {
				log.Fatalf("Could not write to connection: %v", err)
			}
		}(conn)
	}
}
