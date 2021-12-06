package main

import (
	"context"
	blogserver "go-blog-service/blog_server"
)

func main() {
	ch := make(chan int)
	ctx := context.Background()
	go func(channel chan int) {
		blogserver.StartServer(ctx)
	}(ch)

	<-ch
}
