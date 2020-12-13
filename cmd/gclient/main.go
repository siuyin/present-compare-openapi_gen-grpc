package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/siuyin/dflt"
	"google.golang.org/grpc"

	pb "github.com/siuyin/present-compare-openapi_gen-grpc/garith/proto"
)

func main() {
	// 10_OMIT
	// Set up a connection to the server.
	address := dflt.EnvString("ARITH", "localhost:50051")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewArithClient(conn) // HL

	// Make the Sum service call.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Sum(ctx, &pb.SumRequest{A: 2, B: 3}) // HL
	if err != nil {
		log.Fatalf("could not get sum: %v", err)
	}
	fmt.Printf("Sum 2+3 = %v\n", r.GetSum())
	// 20_OMIT
}
