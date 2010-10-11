
//                                 Isolation


package main

import (
	"beanstalk/protocol"
	"net"
)

const (
	nWorkers        = 5
	nJobs           = 1000
	defaultTimeout  = 10 // sec
)

type req struct {
	timeout int
	reply   chan string
}

type client chan req

func (c client) connect(addr string) {
	conn, _ := net.Dial("tcp", addr) // error handling omitted
	for {
		r := <-c
		body := protocol.Reserve(conn, r.timeout)
		r.reply <- body
	}
}

func (c client) reserve(timeout int) string {
	reply := make(chan string)
	c <- req{timeout, reply}
	return <-reply
}

func (c client) work(done chan int) {
	for i := 0; i < nJobs; i++ {
		body := c.reserve(defaultTimeout)
		process(body)
	}
	done <- 1
}

func main() {
	c := make(client)
	go c.connect("localhost:11300")

	done := make(chan int)
	for i := 0; i < nWorkers; i++ {
		go c.work(done)
	}
	for i := 0; i < nWorkers; i++ {
		<-done
	}
}
