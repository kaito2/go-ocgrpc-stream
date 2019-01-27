package main

import (
	"cloud.google.com/go/trace/apiv1"
	"context"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/lovoo/gcloud-opentracing"
	"github.com/opentracing/basictracer-go"
	"github.com/opentracing/opentracing-go"
	pb "go-ocgrpc-stream-try/protos"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	listenAddress = "0.0.0.0:50051"
)

type MySampleStreamServer struct {}

func (s *MySampleStreamServer) GetSampleResponse(req *pb.SampleRequest, stream pb.SampleStream_GetSampleResponseServer) error {
	// 固定の文字列を返すだけ
	res := pb.SampleResponse{Message: "Reply Message"}
	for {
		ctx := stream.Context()
		log.Print(ctx)
		span, _ := opentracing.StartSpanFromContext(ctx, "opentracing_sample.server")
		time.Sleep(1 * time.Second)
		if err := stream.Send(&res); err != nil {
			return err
		}
		span.Finish()
	}
	return nil
}

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

	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.StreamInterceptor(otgrpc.OpenTracingStreamServerInterceptor(tracer)))
	pb.RegisterSampleStreamServer(grpcServer, &MySampleStreamServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}