package main

import (
	"context"
	product "seckill_system/srv-product/kitex_gen/product"
	"seckill_system/srv-product/service"
)

// ProductServiceImpl implements the last service interface defined in the IDL.
type ProductServiceImpl struct{}

// GetProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProduct(ctx context.Context, p *product.Product) (resp *product.Product, err error) {
	// TODO: Your code here...
	_, err = service.GetProduct(p)
	return
}
