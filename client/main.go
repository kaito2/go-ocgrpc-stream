package main

import (
	"cloud.google.com/go/trace/apiv1"
	"context"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/lovoo/gcloud-opentracing"
	"github.com/opentracing/basictracer-go"
	"github.com/opentracing/opentracing-go"
	"go-ocgrpc-stream-try/protos"
	"google.golang.org/grpc"
	"io"
	"log"
)

const (
	//address = "localhost:50051"
	address = "34.73.89.238:50051"
)

func main() {
	// setup open tracing
	ocClient, err := trace.NewClient(context.Background() /*auth options here if necessary*/)
	if err != nil {
		log.Fatalf("error creating a tracing client: %v", err)
	}

	recorder, err := gcloudtracer.NewRecorder(context.Background(), "sansigma-infra", gcloudtracer.TraceClient(ocClient))
	if err != nil {
		log.Fatalf("error creating a recorder: %v", err)
	}
	defer recorder.Close()
	tracer := basictracer.New(recorder)
	opentracing.InitGlobalTracer(tracer)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithStreamInterceptor(otgrpc.OpenTracingStreamClientInterceptor(tracer)))
	if err != nil {
		log.Fatal(err)
	}

	client := grpcstream.NewSampleStreamClient(conn)

	req := grpcstream.SampleRequest{Message: "Request Message"}
	stream, err := client.GetSampleResponse(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}
	for {
		ctx := stream.Context()
		log.Print(ctx)
		span, _ := opentracing.StartSpanFromContext(ctx, "opentracing_sample")
		feature, err := stream.Recv()
		if err == io.EOF { // サーバ側でストリーミングが正常に終了(return nil)された
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Println(feature)
		span.Finish()
	}
}
