package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	myhttp "github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/prometheus/common/log"
	"go-micro-demo/models"
	"io/ioutil"
	"net/http"
)

func callAPI2(s selector.Selector) {
	myClient := myhttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"))
	//req:=myClient.NewRequest("prodservice", "/v1/prods",
	//	map[string]string{})
	//fmt.Println(req.Endpoint())
	//var rsp map[string]interface{}
	//err :=myClient.Call(context.Background(), req, &rsp)
	//if err != nil{
	//	log.Fatal(err)
	//}
	//fmt.Println(rsp["data"])
	req:=myClient.NewRequest("prodservice", "/v1/prods",
		models.ProdsRequest{Size:3})
	var rsp models.ProdListResponse
	err :=myClient.Call(context.Background(), req, &rsp)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(rsp.GetData())
}
func callAPI(addr string, path string, method string)(string, error) {
	req ,_ := http.NewRequest(method, "http://"+addr+path, nil)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err!=nil {
		return "",err
	}
	defer res.Body.Close()

	buf,_ := ioutil.ReadAll(res.Body)
	return string(buf),nil
}

func main() {
	consulReg := consul.NewRegistry(registry.Addrs("192.168.123.217:8500"))
	mySelector := selector.NewSelector(
		selector.Registry(consulReg),
		selector.SetStrategy(selector.RoundRobin),
		)

	//getService, err := consulReg.GetService("prodservice")
	//
	//if err != nil{
	//	log.Fatal(err)
	//}
	//// next := selector.Random(getService)
	//next := selector.RoundRobin(getService)
	//nd,_ :=next()
	//fmt.Println(nd.Id, nd.Address, nd.Metadata)
	//callRes,err := callAPI(nd.Address,"/v1/prods", "GET")
	//if err!= nil{
	//	log.Fatal(err)
	//}
	//fmt.Println(callRes)

	callAPI2(mySelector)
}
