package main

import (
	"context"
	"go-ocgrpc-stream-try/protos"
	"google.golang.org/grpc"
	"io"
	"log"
	trace "cloud.google.com/go/trace/apiv1"
	gcloudtracer "github.com/lovoo/gcloud-opentracing"
	opentracing "github.com/opentracing/opentracing-go"
	basictracer "github.com/opentracing/basictracer-go"
)

const (
	address = "localhost:50051"
)

func main() {
	// setup open tracing
	ocClient, err := trace.NewClient(context.Background() /*auth options here if necessary*/)
	if err != nil {
		log.Fatalf("error creating a tracing client: %v", err)
	}

	recorder, err := gcloudtracer.NewRecorder(context.Background(), "sansigma-infra", ocClient)
	if err != nil {
		log.Fatalf("error creating a recorder: %v", err)
	}
	defer recorder.Close()
	opentracing.InitGlobalTracer(basictracer.New(recorder))

	conn, err := grpc.Dial(address, grpc.WithInsecure())
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
		span, _ := opentracing.StartSpanFromContext(context.Background(), "opentracing_sample")
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
