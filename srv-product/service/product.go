package service

import (
	"context"
	"fmt"
	"log"
	"seckill_system/srv-product/dao"
	"seckill_system/srv-product/kitex_gen/product"
	"seckill_system/srv-product/pdc_model"
	"seckill_system/srv-product/product_idl"
)

func GetProduct(p *product.Product) (info *product.Product, err error) {
	fmt.Println("getProduct id :", p.Id)
	Pdc, err := dao.SearchProductByID(p.Id)
	log.Println("GetProduct test data:", Pdc)
	if err != nil {
		log.Println("GetProduct err:", err)
		return
	}

	return &product.Product{
		Id:          int64(Pdc.ID),
		Name:        Pdc.ProductName,
		Price:       &Pdc.Price,
		Stock:       &Pdc.Stock,
		Image:       &Pdc.Image,
		Description: &Pdc.Description,
	}, nil
}

func GetInfo(p pdc_model.Product) (resp *product.Product, err error) {
	fmt.Println("GetInfo id:", p.ID)

	// 调用rpc
	resp, err = product_idl.ProductClient.GetProduct(context.Background(), &product.Product{
		Id:          int64(p.ID),
		Name:        p.ProductName,
		Price:       &p.Price,
		Stock:       &p.Stock,
		Image:       &p.Image,
		Description: &p.Description,
	})
	if err != nil {
		log.Println("GetInfo rpc call err:", err)
	}
	return
}
