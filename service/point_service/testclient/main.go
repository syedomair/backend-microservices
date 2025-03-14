package main

import (
	"context"
	"log"
	"time"

	pb "github.com/syedomair/backend-microservices/proto/v1/point"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	//conn, err := grpc.NewClient("172.17.0.1:5003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("127.0.0.1:5003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPointServerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUserPoints(ctx, &pb.PointRequest{UserId: "277b1407-738d-408d-91d4-e6a0ead998b1"})
	if err != nil {
		log.Fatalf("could not get Points: %v", err)
	}
	log.Printf("User Point: %s", r.GetUserPoint())
}
