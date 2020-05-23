package Router

import (
	"github.com/gin-gonic/gin"
	"go-micro/Services"
)

func NewRouter(prodService Services.ProdService) *gin.Engine{

	ginRouter := gin.Default()
	ginRouter.Use(InitMiddleware(prodService))
	v1Group :=ginRouter.Group("v1")
	v1Group.Handle("POST", "/prods", Myhandler)
	return ginRouter
}
