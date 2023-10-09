package main

import (
	"context"
	"fmt"
	"github.com/ellioht/coffeeshop/internal/server"
	"os"
	"os/signal"
)

func main() {
	svr := &server.Server{}

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cancel()
	}()

	if err := svr.Run(ctx); err != nil {
		fmt.Println(err)
	}
}
