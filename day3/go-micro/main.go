package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"go-micro/Router"
	"go-micro/Services"
)


type logWrapper struct{
	client.Client
}

func (this *logWrapper)Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", ctx, req.Service(), req.Endpoint())
	return this.Client.Call(ctx, req, rsp)
}

func NewLogWrapper(c client.Client) client.Client {
	return &logWrapper{Client:c}
}
func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.1.56:8500"))


	myService := micro.NewService(micro.Name("prodservice.client"), micro.WrapClient(NewLogWrapper))
	prodService := Services.NewProdService("prodservice",myService.Client())

	httpServer := web.NewService(
		web.Name("httpprodservice"),
		web.Address(":8001"),
		web.Handler(Router.NewRouter(prodService)),
		web.Registry(consulReg))




	httpServer.Init()

	httpServer.Run()
}