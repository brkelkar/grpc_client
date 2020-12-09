package grpcclient

import (
	"context"
	"log"

	"github.com/brkelkar/grpc_client/auth"
	"google.golang.org/grpc"
)

var (
	Conn *grpc.ClientConn
	err  error
)

func Init() {

	Conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer Conn.Close()
}

func AuthMiddleware() {

	c := auth.NewAuthValidationServiceClient(Conn)

	response, err := c.Validate(context.Background(), &auth.Request{AuthToken: "Hello From Client!", UserName: "Bala"})
	if err != nil {
		log.Fatalf("Error when calling Validate: %s", err)
	}
	log.Printf("Response from server: %v", response.Valid)

}
