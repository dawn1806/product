package handler

import (
	"context"
	"fmt"
	"github.com/dawn1806/product/common"
	"github.com/dawn1806/product/domain/model"
	"github.com/dawn1806/product/domain/service"
	. "github.com/dawn1806/product/proto/product"
)
type Product struct{
     ProductDataService service.IProductDataService
}

func (p *Product) AddProduct(ctx context.Context, req *ProductInfo, res *IDResponse) error {
	fmt.Println(req)
	productAdd := &model.Product{}
	if err := common.SwapTo(req, productAdd); err != nil {
		return err
	}
	fmt.Println(productAdd)
	productId, err := p.ProductDataService.AddProduct(productAdd)
	if err != nil {
		return err
	}
	fmt.Println(productId)
	res.ProductId = productId
	return nil
}

func (p *Product) FindProductByID(ctx context.Context, req *IDRequest, res *ProductInfo) error {
	product, err := p.ProductDataService.FindProductByID(req.ProductId)
	if err != nil {
		return err
	}
	if err := common.SwapTo(product, res); err != nil {
		return err
	}
	return nil
}

func (p *Product) UpdateProduct(ctx context.Context, req *ProductInfo, res *ProductResponse) error {
	uProduct := &model.Product{}
	if err := common.SwapTo(req, uProduct); err != nil {
		return err
	}
	if err := p.ProductDataService.UpdateProduct(uProduct); err != nil {
		return err
	}
	res.Message = "商品更新成功"
	return nil
}
func (p *Product) DeleteProductByID(ctx context.Context, req *IDRequest, res *ProductResponse) error {
	if err := p.ProductDataService.DeleteProduct(req.ProductId); err != nil {
		return err
	}

	res.Message = "商品删除成功"
	return nil
}
func (p *Product) FindAllProduct(ctx context.Context, req *AllRequest, res *AllResponse) error {
	productSlice, err := p.ProductDataService.FindAllProduct()
	if err != nil {
		return err
	}

	for _, product := range productSlice {
		productInfo := &ProductInfo{}
		if err := common.SwapTo(product, productInfo); err != nil {
			return err
		}
		res.ProductInfo = append(res.ProductInfo, productInfo)
	}
	return nil
}