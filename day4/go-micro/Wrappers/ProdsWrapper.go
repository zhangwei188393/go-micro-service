package Wrappers

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
	"go-micro/Services"
	"strconv"
)

type ProdsWrapper struct{
	client.Client
}

func newProd(id int32, pname string)*Services.ProdModel{
	return &Services.ProdModel{ProdID:id, ProdName:pname}
}

func defaultProds(rsp interface{}) {
	models := make([]*Services.ProdModel,0)
	var i int32
	for i=0;i<4; i++ {
		models=append(models, newProd(50+i, "prodname"+ strconv.Itoa(50+int(i))))
	}
	result := rsp.(*Services.ProdListResponse)
	result.Data = models
}

func (this *ProdsWrapper)Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	//fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", ctx, req.Service(), req.Endpoint())

	//return this.Client.Call(ctx, req, rsp)
	cmdName := req.Service()+"."+req.Endpoint()
	configA := hystrix.CommandConfig{Timeout:100}
	hystrix.ConfigureCommand(cmdName, configA)
	return hystrix.Do(cmdName, func() error {
		return this.Client.Call(ctx, req, rsp)
	}, func(e error) error {
		defaultProds(rsp)
		return nil
	})
}

func NewProdsWrapper(c client.Client) client.Client {
	return &ProdsWrapper{c}
}