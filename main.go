package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()

	conn, _, err := DialWithContext(ctx, "tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		time.Sleep(1 * time.Second)

		_, err = conn.Write([]byte("Hello World!!"))

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("書き込み")
	}

}

func DialWithContext(ctx context.Context, network, address string) (net.Conn, func() bool, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, nil, err
	}

	stop := context.AfterFunc(ctx, func() {
		fmt.Println("アフターファンク")
		conn.Close()
	})

	return conn, stop, nil
}
