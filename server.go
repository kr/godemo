
//                                   Server


package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

var queue = make(chan string)

func Serve(c net.Conn) {
	defer c.Close()
	defer fmt.Println("closing", c.RemoteAddr())
	fmt.Println("serving", c.RemoteAddr())

	r := bufio.NewReader(c)

	for {
		var reply string

		line, err := r.ReadString('\n')
		if err != nil { // e.g. EOF
			return
		}


		switch {
		case strings.HasPrefix(line, "put "):
			go func(s string) {
				queue <- s // blocks
			}(line[4:])
			reply = "ok\r\n"

		case line == "get\r\n":
			line = <-queue // blocks
			reply = "ok " + line

        default:
			reply = "unknown command\r\n"
		}


		for len(reply) > 0 {
			n, err := io.WriteString(c, reply)
			if err != nil { // e.g. EOF
				return
			}
			reply = reply[n:]
		}
	}
}

func main() {
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go Serve(c)
	}
}
