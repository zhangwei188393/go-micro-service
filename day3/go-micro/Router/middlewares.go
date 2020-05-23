package Router

import (
	"github.com/gin-gonic/gin"
	"go-micro/Services"
)

func InitMiddleware(prodService Services.ProdService) gin.HandlerFunc{
	return  func(context *gin.Context) {
		context.Keys = make(map[string]interface{})
		context.Keys["prodservice"] = prodService
		context.Next()
	}
}
