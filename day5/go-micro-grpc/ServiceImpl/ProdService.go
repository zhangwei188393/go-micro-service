package ServiceImpl

import (
	"context"
	"go-micro-grpc/Services"
	"strconv"
	"time"
)

type ProdService struct {

}

func newProd(id int32, pname string)*Services.ProdModel{
	return &Services.ProdModel{ProdID:id, ProdName:pname}
}
func(*ProdService) GetProdsList(ctx context.Context, in *Services.ProdsRequest, res *Services.ProdListResponse) error{
	time.Sleep(3*time.Second)
	models := make([]*Services.ProdModel,0)
	var i int32
	for i=0;i<in.Size; i++ {
		models=append(models, newProd(100+i, "prodname"+ strconv.Itoa(100+int(i))))
	}
	res.Data = models
	return nil
}

func (*ProdService) GetProdDetail(ctx context.Context, in *Services.ProdsRequest, res *Services.ProdDetailResponse) error {
	res.Data = newProd(in.ProdId, "single product")
	return nil
}

