package main

import (
	"io"
	"log"
	"net"
	"os"
)

func response(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	conn.Write([]byte("hello tcp echo"))
	response(os.Stdout, conn)
}
