package main

import (
	"context"
	"fmt"
	"log"

	sw "github.com/siuyin/present-compare-openapi_gen-grpc/client"
)

func main() {
	// 10_OMIT
	fmt.Println("arith API client")
	cl := sw.NewAPIClient(sw.NewConfiguration())
	sum, _, err := cl.ArithmeticApi.SumGet(context.TODO(), 2, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("sum of %d and %d is %v\n", 2, 3, sum.Sum)
	// 20_OMIT
}
