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

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	oid := result.InsertedID.(primitive.ObjectID)
	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       oid.String(),
			AuthorId: blog.GetAuthorId(),
			Title:    blog.GetTitle(),
			Content:  blog.GetContent(),
		},
	}, nil
}

func (*server) ReadBlog(ctx context.Context, req *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	blogId, err := primitive.ObjectIDFromHex(req.GetBlogId())
	errorChecker.HasError(err).Fatal("Invalid argument")

	data := &models.Blog{}
	filter := bson.D{{"_id", blogId}}
	findErr := mongoclient.BlogCollection.FindOne(context.Background(), filter).Decode(data)
	errorChecker.HasError(findErr).Info("Couldn't find blog with given BlogId")

	return &blogpb.ReadBlogResponse{
		Blog: &blogpb.Blog{
			Id:       data.Id.String(),
			AuthorId: data.AuthorId,
			Title:    data.Title,
			Content:  data.Content,
		},
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
