package main

import (
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
		time.Sleep(1 * time.Second)
		if err := stream.Send(&res); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSampleStreamServer(grpcServer, &MySampleStreamServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}