package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ArpitKRS/go-grpc/proto"
)

func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client streaming is started")
	stream, err := client.SayHelloCLientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send the names: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Sent the request with name: %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished")
	if err != nil {
		log.Fatalf("Error while receiving %v", err)
	}
	log.Printf("%v", res.Messages)
}