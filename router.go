package main

import (
	"github.com/hhy5861/gin-base/controllers"
	"github.com/hhy5861/gin-base/grpc_service"
	pbAccount "github.com/hhy5861/gin-base/proto"
	"github.com/hhy5861/go-common/gin_controller"
	"google.golang.org/grpc"
)

func (svc *Server) Router() {
	ctl := &gin_controller.Controller{
		JwtConf: svc.Config.Security.Jwt,
	}

	accountController := controllers.NewAccountControllers(ctl)

	ignoreAuthGroup := svc.BaseServer.RouterGroup.Group("/v1/common")
	{
		ignoreAuthGroup.POST("/account/login", accountController.Login)
		ignoreAuthGroup.POST("/account/grpc/login", accountController.GrpcLogin)
	}

	authGroup := svc.BaseServer.RouterGroup.Group("/v1", gin_controller.NewController(ctl).JwtInterceptor)
	{
		authGroup.GET("/account/list", accountController.List)
	}
}

func (svc *Server) RegisteGrpcRouter(grpcSvc *grpc.Server) {
	pbAccount.RegisterAccountServer(svc.BaseServer.GrpcServer.Server, grpc_service.NewAccountGrpcService())
}
