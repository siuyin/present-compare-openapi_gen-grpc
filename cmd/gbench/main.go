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
	c := pb.NewArithClient(conn)

	start := time.Now()
	for i := 0; i < 2000; i++ {
		callSum(c)
	}
	end := time.Now()
	dur := end.Sub(start).Seconds()
	fmt.Printf("elapsed time: %v (%v requests/second)\n", dur, 2000.0/dur)
	// 20_OMIT
}
func callSum(c pb.ArithClient) {
	// Make the Sum service call.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.Sum(ctx, &pb.SumRequest{A: 2, B: 3}) // HL
	if err != nil {
		log.Fatalf("could not get sum: %v", err)
	}
}
