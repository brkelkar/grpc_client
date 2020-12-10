package grpcclient

import (
	"context"
	"log"

	"github.com/brkelkar/common_utils/logger"
	"github.com/brkelkar/grpc_client/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var (
	Conn *grpc.ClientConn
	err  error
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func GetConnectionID() *grpc.ClientConn {
	Conn, err = grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	return Conn
}

//AuthMiddleware is to authenticate user
func AuthMiddleware(Conn *grpc.ClientConn) gin.HandlerFunc {

	authClient := auth.NewAuthValidationServiceClient(Conn)

	response, err := authClient.Validate(context.Background(), &auth.Request{AuthToken: "Hello From Client!", UserName: "Bala"})
	if err != nil {
		logger.Error("Error when calling Validate: %s", err)
	}

	if response.Valid == true {
		return func(c *gin.Context) {
			c.Next()
		}
	}

	return func(c *gin.Context) {
		respondWithError(c, 401, "Invalid API token")
		c.Next()

	}

}
