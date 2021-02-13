package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/dawn1806/product/common"
	go_micro_service_product "github.com/dawn1806/product/proto/product"
)

func main() {
	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer("go.micro.service.product.client", "localhost:6831")
	if err != nil {
		fmt.Println("[client main] common.NewTracer err=", err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	service := micro.NewService(
		micro.Name("go.micro.service.product.client"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8085"),
		micro.Registry(consulRegistry),
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		)

	productService := go_micro_service_product.NewProductService("go.micro.service.product", service.Client())

	productAdd := &go_micro_service_product.ProductInfo{
		ProductName:        "name01",
		ProductSku:         "sku1",
		ProductPrice:       199,
		ProductDescription: "hehehhhhhhh",
		ProductCategoryId:  1,
		ProductImage:       []*go_micro_service_product.ProductImage{
			{
				ImageName: "img001",
				ImageCode: "000999",
				ImageUrl:  "xxxxx",
			},
			{
				ImageName: "img002",
				ImageCode: "4j4j4j4j4",
				ImageUrl:  "xxxxx",
			},
		},
		ProductSize:        []*go_micro_service_product.ProductSize{
			{
				SizeName: "middle",
				SizeCode: "hehhhhhh",
			},
		},
		ProductSeo:         &go_micro_service_product.ProductSeo{
			SeoTitle:       "title001",
			SeoKeywords:    "slg",
			SeoDescription: "thi shis  sdkdlgs   sdgsg",
			SeoCode:        "qqqq",
		},
	}

	res, err := productService.AddProduct(context.TODO(), productAdd)
	if err != nil {
		fmt.Println("[client main] AddProduct err=", err)
	}
	fmt.Println("[client main] add success. 产品ID：", res.ProductId)
}
