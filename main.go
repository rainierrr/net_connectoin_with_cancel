package main

import (
	"context"
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
			panic(err)
		}
	}
}
