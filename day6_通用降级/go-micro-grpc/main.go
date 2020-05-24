package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"

	"github.com/micro/go-plugins/registry/consul"
	"go-micro-grpc/ServiceImpl"
	"go-micro-grpc/Services"
)
func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.1.56:8500"))

	prodService := micro.NewService(
		micro.Name("prodservice"),
		micro.Address(":8002"),
		micro.Registry(consulReg))

	prodService.Init()
	Services.RegisterProdServiceHandler(prodService.Server(), new(ServiceImpl.ProdService))
	prodService.Run()
}