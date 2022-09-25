package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-grpc-gateway/pkg/auth/pb"
	"net/http"
	"strings"
)

type MiddleWareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) MiddleWareConfig {
	return MiddleWareConfig{svc}
}

func (c *MiddleWareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.UserId)
	ctx.Next()
}
