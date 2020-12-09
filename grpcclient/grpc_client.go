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

//AuthMiddleware is to authenticate user
func AuthMiddleware(c *gin.Context) {

	authClient := auth.NewAuthValidationServiceClient(Conn)

	response, err := authClient.Validate(context.Background(), &auth.Request{AuthToken: "Hello From Client!", UserName: "Bala"})
	if err != nil {
		logger.Error("Error when calling Validate: %s", err)
	}

	if response.Valid == true {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusForbidden)
	}

}
