package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/siuyin/dflt"
	"google.golang.org/grpc"

	pb "github.com/siuyin/present-compare-openapi_gen-grpc/garith/proto"
)

// 10_OMIT
type server struct {
	pb.UnimplementedArithServer // defined in garith/proto/arith_grpc.pb.go
}

// Sum adds two numbers given in SumRequest.
func (s *server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum received: %v, %v", in.GetA(), in.GetB())
	return &pb.SumResponse{Sum: in.GetA() + in.GetB()}, nil
}

// 20_OMIT
func main() {
	fmt.Println("gRPC Arith server")
	port := dflt.EnvString("PORT", "50051")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterArithServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// 30_OMIT
