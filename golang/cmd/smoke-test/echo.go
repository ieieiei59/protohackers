package main

import (
	"fmt"
	"log"
	"net"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func receiveTCPConnection(listener *net.TCPListener) {
	for {
		log.Println("Connection Start")

		conn, err := listener.AcceptTCP()
		if err != nil {
			switch err := err.(type) {
			case net.Error:
				if err.Timeout() {
					log.Println("Connection Close")
					return
				}
			default:
				log.Println("Another Error")
				return
			}
		}

		go func(conn *net.TCPConn) {
			echoHandler(conn)
		}(conn)
	}
}

func echoHandler(conn *net.TCPConn) {
	defer conn.Close()

	b := make([]byte, 1024)
	n, _ := conn.Read(b)
	conn.Write(b[:n])

	log.Println(fmt.Sprintf("readed: `%s`", string(b[:n])))
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080")
	logFatal(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	logFatal(err)

	fmt.Println("Start TCP Server...")
	receiveTCPConnection(listener)
}
