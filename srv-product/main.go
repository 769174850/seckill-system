package main

import (
	"log"
	product "seckill_system/srv-product/kitex_gen/product/productservice"
)

func main() {
	svr := product.NewServer(new(ProductServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
