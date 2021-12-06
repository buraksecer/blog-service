package blogserver

import (
	"fmt"
	"go-blog-service/blogpb"
	"go-blog-service/utils"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

const (
	port = 50051
)

type server struct{}

var errorChecker utils.ErrorChecker

func StartServer() {
	errorChecker = utils.NewErrorChecker()
	fmt.Printf("Starting Blog Server...\nListening port %v \n", port)
	listener, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(port))

	errorChecker.HasError(err).Fatal("")

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, &server{})

	serveErr := s.Serve(listener)

	errorChecker.HasError(serveErr).Fatal("")
}
