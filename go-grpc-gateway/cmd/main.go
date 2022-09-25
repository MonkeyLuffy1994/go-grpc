package main

import (
	"github.com/gin-gonic/gin"
	"go-grpc-gateway/pkg/auth"
	"go-grpc-gateway/pkg/config"
	"go-grpc-gateway/pkg/order"
	"go-grpc-gateway/pkg/product"
	"log"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterOrderRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
