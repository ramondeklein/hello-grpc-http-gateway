package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/ramondeklein/grpc-json/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcPort = flag.Int("grpc-port", 9090, "gRPC port")
	httpPort = flag.Int("http-port", 8080, "HTTP port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()

	errCh := make(chan error)

	// gRPC server
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
		if err != nil {
			errCh <- err
			return
		}

		s := grpc.NewServer()
		pb.RegisterGreeterServer(s, &server{})
		log.Printf("gRPC server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			errCh <- err
			return
		}
	}()

	// HTTP server
	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *httpPort))
		if err != nil {
			errCh <- err
			return
		}

		// Register gRPC server endpoint
		// Note: Make sure the gRPC server is running properly and accessible
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		err = pb.RegisterGreeterHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%d", *grpcPort), opts)
		if err != nil {
			errCh <- err
			return
		}

		// Start HTTP server (and proxy calls to gRPC server endpoint)
		log.Printf("gRPC server listening at %v", lis.Addr())
		if err := http.Serve(lis, mux); err != nil {
			errCh <- err
			return
		}
	}()

	<-errCh
}
