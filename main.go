package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	conn, err := net.Dial(listener.Addr().Network(), listener.Addr().String())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// ここでタイムアウト
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	newConn, err := DialWithContext(ctx, listener.Addr().Network(), listener.Addr().String())
	if err != nil {
		fmt.Println(err)
		return
	}
	//スコープから出るときに
	defer newConn.Close()
	fmt.Println(err)

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
