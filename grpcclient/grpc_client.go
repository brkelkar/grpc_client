package grpcclient

import (
	"context"
	"net/http"

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

//AuthMiddleware is to authenticate user
func AuthMiddleware() gin.HandlerFunc {

	authClient := auth.NewAuthValidationServiceClient(Conn)

	response, err := authClient.Validate(context.Background(), &auth.Request{AuthToken: "Hello From Client!", UserName: "Bala"})
	if err != nil {
		logger.Error("Error when calling Validate: %s", err)
	}

	if response.Valid == true {
		return func(c *gin.Context) {
			c.Next()
		}
	} else {
		return func(c *gin.Context) {
			respondWithError(c, 401, "Invalid API token")
			c.Next()
	}

}
