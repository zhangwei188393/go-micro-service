package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"go-micro/Services"
)
func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.123.217:8500"))

	ginRouter := gin.Default()



	httpServer := web.NewService(
		web.Name("httpprodservice"),
		web.Address(":8001"),
		web.Handler(ginRouter),
		web.Registry(consulReg))

	myService := micro.NewService(micro.Name("prodservice.client"))
	prodService := Services.NewProdService("prodservice",myService.Client())

	fmt.Printf("%T\n",prodService)
	v1Group :=ginRouter.Group("v1")
	v1Group.Handle("POST", "/prods", func(ginCtx *gin.Context) {
		var prodReq Services.ProdsRequest
		err:=ginCtx.Bind(&prodReq)
		if err!=nil{
			ginCtx.JSON(500, gin.H{"status":err.Error()})
		} else {
			prodRes,_ := prodService.GetProdsList(context.Background(), &prodReq)
			ginCtx.JSON(200, gin.H{"data":prodRes.GetData()})
		}


	})
	httpServer.Init()

	httpServer.Run()
}