package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tetratelabs/go2sky"
	"github.com/tetratelabs/go2sky/reporter"
)

// ExampleNewTracer test
func ExampleNewTracer() {
	// Use gRPC reporter for production
	// r, err := reporter.NewLogReporter()
	r, err := reporter.NewGRPCReporter("127.0.0.1:11800")
	if err != nil {
		log.Fatalf("new reporter error %v \n", err)
	}
	defer r.Close()
	tracer, err := go2sky.NewTracer("example", go2sky.WithReporter(r))
	if err != nil {
		log.Fatalf("create tracer error %v \n", err)
	}
	// This for test
	tracer.WaitUntilRegister()
	span, ctx, err := tracer.CreateLocalSpan(context.Background())
	if err != nil {
		log.Fatalf("create new local span error %v \n", err)
	}
	span.SetOperationName("invoke data")
	span.Tag("kind", "outer")
	time.Sleep(time.Second)
	subSpan1, _, err := tracer.CreateLocalSpan(ctx)
	if err != nil {
		log.Fatalf("create new sub local span error %v \n", err)
	}
	subSpan1.SetOperationName("invoke inner")
	subSpan1.Log(time.Now(), "inner", "this is right")
	time.Sleep(time.Second)
	subSpan1.End()

	subSpan2, _, err := tracer.CreateLocalSpan(ctx)
	if err != nil {
		log.Fatalf("create new sub local span error %v \n", err)
	}
	subSpan2.SetOperationName("alalalallalala")
	subSpan2.Log(time.Now(), "abcdefg", "wahtwaht?")
	time.Sleep(time.Second)
	subSpan2.End()

	time.Sleep(500 * time.Millisecond)
	span.End()
	time.Sleep(time.Second)
	// Output:
}

func main() {
	ExampleNewTracer()
	fmt.Println("hello world")
}
