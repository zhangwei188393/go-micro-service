package main

//import (
//	"github.com/micro/go-micro/web"
//	"github.com/gin-gonic/gin"
//
//)
//func main() {
//
//	ginRouter := gin.Default()
//	ginRouter.Handle("GET", "/user", func(ctx *gin.Context){
//		ctx.String(200, "user api")
//	})
//	ginRouter.Handle("GET", "/news", func(ctx *gin.Context){
//		ctx.String(200, "news api")
//	})
//
//
//	server := web.NewService(web.Address(":8001"), web.Handler(ginRouter))
//
//	//server.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
//	//	writer.Write([]byte("hello World"))
//	//})
//	server.Run()
//}
/////////////////////////////////////////

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"go-micro-demo/ProdService"
	"go-micro-demo/helper"
)
func main() {
	consulReg :=consul.NewRegistry(
		registry.Addrs("192.168.123.217:8500"))
	ginRouter := gin.Default()
	//ginRouter.Handle("GET", "/user", func(ctx *gin.Context){
	//	ctx.String(200, "user api")
	//})
	//ginRouter.Handle("GET", "/news", func(ctx *gin.Context){
	//	ctx.String(200, "news api")
	//})

	//ginRouter.Handle("GET", "/", func(context *gin.Context){
	//	data :=make([]interface{},0)
	//	context.JSON(200, gin.H{
	//		"data":data,
	//	})
	//})
	v1 :=ginRouter.Group("v1")
	v1.Handle("POST", "/prods", func(context *gin.Context) {
		var pr helper.ProdRequest
		err:=context.Bind(&pr)
		if err!= nil || pr.Size <= 0 {
			pr = helper.ProdRequest{Size:2}
		}
		fmt.Println(pr.Size)
		context.JSON(200, gin.H{
			"data":ProdService.NewProdList(pr.Size),
		})
	})

	server := web.NewService(
		web.Address(":8001"),
		web.Handler(ginRouter),
		web.Registry(consulReg),
		web.Name("prodservice"),
	)
	server.Init()
	server.Run()
}
