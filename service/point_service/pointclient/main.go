package main

import (
	"context"
	"log"
	"time"

	pb "github.com/syedomair/backend-microservices/service/point_service/pointserver/point"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient("localhost:8185", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPointServerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUserPoints(ctx, &pb.PointRequest{UserId: "0075dcd1-faaa-4083-b063-c959eff537dd"})
	if err != nil {
		log.Fatalf("could not get Points: %v", err)
	}
	log.Printf("User Point: %s", r.GetUserPoint())
}
