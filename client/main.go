package main

import (
	"context"
	"go-ocgrpc-stream-try/protos"
	"google.golang.org/grpc"
	"io"
	"log"
)

const (
	address = "localhost:50051"
)

func main() {
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
		feature, err := stream.Recv()
		if err == io.EOF { // サーバ側でストリーミングが正常に終了(return nil)された
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Println(feature)
	}
}
