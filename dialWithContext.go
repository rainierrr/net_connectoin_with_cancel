package main

import (
	"context"
	"net"
)

func DialWithContext(ctx context.Context, network, address string) (net.Conn, func() bool, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, nil, err
	}

	stop := context.AfterFunc(ctx, func() {
		conn.Close()
	})

	return conn, stop, nil
}
