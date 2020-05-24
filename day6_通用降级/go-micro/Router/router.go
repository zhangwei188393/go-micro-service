package Router

import (
	"github.com/gin-gonic/gin"
	"go-micro/Services"
)

func NewRouter(prodService Services.ProdService) *gin.Engine{

	ginRouter := gin.Default()
	ginRouter.Use(InitMiddleware(prodService),ErrorMiddleware())
	v1Group :=ginRouter.Group("v1")
	v1Group.Handle("POST", "/prods", GetProdList)
	v1Group.Handle("GET", "/prods/:pid", GetProdDetail)

	return ginRouter
}
