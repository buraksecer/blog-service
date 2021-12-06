package main

import (
	"context"
	blogclient "go-blog-service/blog_client"
	blogserver "go-blog-service/blog_server"
	"os"
	"os/signal"
)

func main() {
	ch := make(chan os.Signal, 1)
	ctx := context.TODO()
	go func(c context.Context) {
		blogserver.StartServer(c)
	}(ctx)

	blogclient.StartClient()
	signal.Notify(ch, os.Interrupt)
	<-ch
}
