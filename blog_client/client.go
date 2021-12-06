package blogclient

import (
	"context"
	"go-blog-service/blogpb"
	"go-blog-service/utils"
	"log"

	"google.golang.org/grpc"
)

var errChecker utils.ErrorChecker

func StartClient() {
	log.Println("Blog Client started...")

	opts := grpc.WithInsecure()
	conn, err := grpc.Dial("localhost:50051", opts)
	defer conn.Close()
	errChecker.HasError(err).Fatal("An error occured when starting blog client")

	client := blogpb.NewBlogServiceClient(conn)

	client.CreateBlog(context.TODO(), &blogpb.CreateBlogRequest{
		Blog: &blogpb.Blog{
			AuthorId: "123-444",
			Title:    "About Go",
			Content:  "Go is the best !",
		},
	})
}
