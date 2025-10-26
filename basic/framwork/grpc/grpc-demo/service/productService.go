package service

import (
	"context"
	"log"
)

// 自定义服务结构体，必须嵌入 UnimplementedProductServiceServer 以保证兼容性
type ProductService struct {
	UnimplementedProductServiceServer // 嵌入未实现的基础结构
}

// 实现 GetProductStock 方法（对应 proto 中定义的服务方法）
func (s *ProductService) GetProductStock(ctx context.Context, req *ProductRequest) (*ProductResponse, error) {
	// 这里写业务逻辑，示例：根据请求的商品ID返回库存
	log.Printf("收到查询商品 %d 库存的请求", req.ProductId)
	// 假设库存为100
	return &ProductResponse{ProductStock: 100}, nil
}
