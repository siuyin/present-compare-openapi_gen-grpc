package main

import (
	"context"
	"fmt"
	"log"

	sw "github.com/siuyin/present-compare-openapi_gen-grpc/client"
)

func main() {
	fmt.Println("arith API client")
	cl := sw.NewAPIClient(sw.NewConfiguration())
	sum, _, err := cl.ArithmeticApi.SumGet(context.Background(), 2, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("sum of %d and %d is %v\n", 2, 3, sum.Sum)
}
