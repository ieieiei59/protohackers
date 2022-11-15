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
	
	for {
		b := make([]byte, 16)	
		n, err := conn.Read(b)
		
		if err != nil {
			log.Println("close")
			break
		}

		conn.Write(b[:n])

		log.Println(fmt.Sprintf("readed (%d): `%s`", n, string(b[:n])))
	}
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8080")
	logFatal(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	logFatal(err)

	fmt.Println("Start TCP Server...")
	receiveTCPConnection(listener)
}
