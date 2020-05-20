package ProdService

import (
	"fmt"
	"strconv"
)

//type ProdModel struct {
//	ProdID int
//	ProdName string
//}


type ProdModel struct {
	ProdID int		`json:"pid"`
	ProdName string		`json:"pname"`
}

func NewProd(id int,  pname string) *ProdModel {
	return &ProdModel{ProdID:id, ProdName:pname}
}


func NewProdList(n int) []*ProdModel{
	ret:=make([]*ProdModel, 0)
	for i:=0; i<n; i++{
		ret=append(ret, NewProd(i, "ProdName"+ strconv.Itoa(100+i)))
	}
	fmt.Println(ret[0])
	return ret
}