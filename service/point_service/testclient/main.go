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

	//conn, err := grpc.NewClient("127.0.0.1:5003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("127.0.0.1:8185", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPointServerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUserPoints(ctx, &pb.PointRequest{UserId: "520349ca-a915-487d-ba21-c9740dd08aa7"})
	if err != nil {
		log.Fatalf("could not get Points: %v", err)
	}
	userIDs := []string{"ee419aab-2b95-48fd-a871-7bc3eb8d244e", "dc78e441-ab03-413d-a485-a7325b271147", "520349ca-a915-487d-ba21-c9740dd08aa7"}
	r1, err := c.GetUserListPoints(ctx, &pb.UserListRequest{UserIds: userIDs})
	if err != nil {
		log.Fatalf("could not get Points: %v", err)
	}
	log.Printf("User Point: %s", r.GetUserPoint())
	log.Printf("User Point: %v", r1.GetUserPoints())
}
