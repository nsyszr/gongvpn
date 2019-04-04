package main

import (
	"io"
	"log"
	"net"
	"os"
	"os/signal"
)

func reader(r io.Reader, bufChan chan []byte) {
	buf := make([]byte, 1024)
	for {

		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		// println("Client got:", string(buf[0:n]))
		bufChan <- buf[0:n]
	}
}

func main() {
	c, err := net.Dial("unix", "/home/tlx3m3j/src/github.com/nsyszr/ngvpn/ovpn_demo/mgmt.sock")
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer c.Close()

	bufChan := make(chan []byte)

	go reader(c, bufChan)
	go func() {
		for {
			select {
			case buf := <-bufChan:
				println("Client got:", string(buf))
			}
		}
	}()

	c.Write([]byte("bytecount 1\n"))

	// Wait for interrupt signal to gracefully shutdown the server
	quitChan := make(chan os.Signal)
	signal.Notify(quitChan, os.Interrupt)
	<-quitChan

	/*for {
		msg := "hi"
		_, err := c.Write([]byte(msg))
		if err != nil {
			log.Fatal("Write error:", err)
			break
		}
		println("Client sent:", msg)
		time.Sleep(1e9)
	}*/

}
