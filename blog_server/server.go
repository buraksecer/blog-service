package blogserver

import (
	"context"
	"fmt"
	mongoclient "go-blog-service/blog_server/mongo_client"
	"go-blog-service/blog_server/mongo_client/models"
	"go-blog-service/blogpb"
	"go-blog-service/utils"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/****** CONSTANTS & UTILS *******/

var errorChecker utils.ErrorChecker

const (
	port = 50051
)

/****** CONSTANTS & UTILS *******/

type server struct{}

func (*server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (res *blogpb.CreateBlogResponse, err error) {
	blog := req.GetBlog()
	data := models.Blog{
		AuthorId: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}

	result, err := mongoclient.BlogCollection.InsertOne(context.TODO(), data)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal server error : %v", err))
	}

	blog.Id = result.InsertedID.(string)
	return &blogpb.CreateBlogResponse{
		Blog: blog,
	}, nil
}

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
