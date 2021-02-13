package main

import (
	"fmt"
	"github.com/dawn1806/product/common"
	"github.com/dawn1806/product/domain/repository"
	service2 "github.com/dawn1806/product/domain/service"
	"github.com/dawn1806/product/handler"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"

	product "github.com/dawn1806/product/proto/product"
)

func main() {

	// 配置中心
	conf, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		fmt.Println("[main] config err=", err)
	}

	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// 链路追踪
	//t, io, err := common.NewTracer("go.micro.service.product", "localhost:6831")
	//if err != nil {
	//	fmt.Println("[main] common.NewTracer err=", err)
	//}
	//defer io.Close()
	//opentracing.SetGlobalTracer(t)

	// 数据库设置
	mysqlInfo := common.GetMysqlFromConsul(conf, "mysql")
	db, err := gorm.Open("mysql", mysqlInfo.User + ":" + mysqlInfo.Password + "@/" + mysqlInfo.Database + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("[main] gorm.Open err=", err)
	}
	defer db.Close()
	db.SingularTable(true)

	//repository.NewProductRepository(db).InitTable()

	iproductService := service2.NewProductDataService(repository.NewProductRepository(db))

	// New Service
	service := micro.NewService(
		micro.Name("micro.product"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8002"),
		micro.Registry(consulRegistry),
		//micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	// Initialise service
	service.Init()

	// Register Handler
	product.RegisterProductHandler(service.Server(), &handler.Product{ProductDataService: iproductService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
