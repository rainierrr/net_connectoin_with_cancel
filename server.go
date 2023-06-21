package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		data := make([]byte, 1024)
		_, err = conn.Read(data)

		if err != nil {
			fmt.Println("cannot write", err)
		}

		fmt.Println(string(data))
	}
}
