package Router

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-micro/Services"
)
//func newProd(id int32, pname string)*Services.ProdModel{
//	return &Services.ProdModel{ProdID:id, ProdName:pname}
//}
//
//func defaultProds() (*Services.ProdListResponse, error) {
//	models := make([]*Services.ProdModel,0)
//	var i int32
//	for i=0;i<4; i++ {
//		models=append(models, newProd(50+i, "prodname"+ strconv.Itoa(50+int(i))))
//	}
//	res := &Services.ProdListResponse{}
//	res.Data = models
//	return res, nil
//}
func Myhandler(ginCtx *gin.Context) {
	var prodReq Services.ProdsRequest

	prodService := ginCtx.Keys["prodservice"].(Services.ProdService)
	err:=ginCtx.Bind(&prodReq)
	if err!=nil{
		ginCtx.JSON(500, gin.H{"status":err.Error()})
	} else {

		var prodRes *Services.ProdListResponse
		prodRes,_ = prodService.GetProdsList(context.Background(), &prodReq)

		ginCtx.JSON(200, gin.H{"data":prodRes.GetData()})

		//configA := hystrix.CommandConfig{Timeout:1000}
		//hystrix.ConfigureCommand("getprods", configA)
		//err:=hystrix.Do("getprods", func() error {
		//	prodRes,err = prodService.GetProdsList(context.Background(), &prodReq)
		//	return err
		//}, func(e error) error {
		//	prodRes,err = defaultProds()
		//	return err
		//})
		//if err!= nil{
		//	ginCtx.JSON(500, gin.H{"status":err.Error()})
		//} else {
		//	ginCtx.JSON(200, gin.H{"data":prodRes.GetData()})
		//}

		//prodRes,_ := prodService.GetProdsList(context.Background(), &prodReq)
		//ginCtx.JSON(200, gin.H{"data":prodRes.GetData()})
	}


}
