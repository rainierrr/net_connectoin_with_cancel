package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

func main() {

	// ここでタイムアウト
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()

	//conn, err := net.Dial("tcp", ":8080")
	conn, err := DialWithContext(ctx, "tcp", ":8080")
	if err != nil {
		panic(err)
	}

	_, err = conn.Write([]byte("Hello World!!"))
	if err != nil {
		panic(err)
	}

	//スコープから出るときに
	defer conn.Close()
}

func DialWithContext(ctx context.Context, network, address string) (net.Conn, error) {

	fmt.Println("実行！")
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	stopc := make(chan struct{})
	stop := context.AfterFunc(ctx, func() {
		fmt.Println("アフターファンク")
		conn.SetDeadline(time.Now())
		close(stopc)
	})

	if !stop() {
		// The AfterFunc was started.
		// Wait for it to complete, and reset the Conn's deadline.
		<-stopc
		conn.SetReadDeadline(time.Time{})
		fmt.Println("リセット")
		return conn, ctx.Err()
	}

	return conn, nil
}
