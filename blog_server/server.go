package blogserver

import (
	"context"
	mongoclient "go-blog-service/blog_server/mongo_client"
	"go-blog-service/blogpb"
	"go-blog-service/utils"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

const (
	port = 50051
)

type server struct{}

var errorChecker utils.ErrorChecker

func StartServer(ctx context.Context) {
	log.Printf("Starting Blog Server... Listening port %v", port)
	mongoclient.StartMongoClient()
	defer mongoclient.CloseConnection()

	listener, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(port))

	errorChecker.HasError(err).Fatal("")

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, &server{})

	serveErr := s.Serve(listener)

	errorChecker.HasError(serveErr).Fatal("")
}
