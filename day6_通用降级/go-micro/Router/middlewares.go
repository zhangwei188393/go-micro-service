package Router

import (
	"fmt"
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

func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r:=recover(); r!=nil{
				ctx.JSON(500,gin.H{"status":fmt.Sprint("%s",r)})
				ctx.Abort()
			}
		}()
		ctx.Next()
	}

}
