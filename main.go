package main

import (
	"context"
	blogserver "go-blog-service/blog_server"
	"os"
	"os/signal"
)

func main() {
	ch := make(chan os.Signal, 1)
	ctx := context.Background()
	go func(c context.Context) {
		blogserver.StartServer(c)
	}(ctx)

	signal.Notify(ch, os.Interrupt)
	<-ch
}
