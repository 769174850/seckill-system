package product_idl

import (
	"github.com/cloudwego/kitex/client"
	"log"
	product "seckill_system/srv-product/kitex_gen/product/productservice"
)

var ProductClient product.Client

func InitProductClient() {
	p, err := product.NewClient("Product", client.WithHostPorts("localhost:8888"))
	if err != nil {
		log.Println("can not connect to product service:", err)
		return
	}
	ProductClient = p
}
