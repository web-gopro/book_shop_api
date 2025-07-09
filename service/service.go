package service

import (
	"fmt"

	"github.com/web-gopro/book_shop_api/genproto/book_shop"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Service() ServiceManagerI {

	userService, err := grpc.NewClient("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	productService, err := grpc.NewClient("localhost:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println(err)
		return nil
	}

	serviseManager := &serviceManager{
		userService:    book_shop.NewUserServiceClient(userService),
		productService: book_shop.NewProductServiceClient(productService),
	}

	return serviseManager
}

type ServiceManagerI interface {
	GetUserSevice() book_shop.UserServiceClient
	GetProductSevice() book_shop.ProductServiceClient
}

type serviceManager struct {
	userService    book_shop.UserServiceClient
	productService book_shop.ProductServiceClient
}

func (s *serviceManager) GetUserSevice() book_shop.UserServiceClient {

	return s.userService
}

func (s *serviceManager) GetProductSevice() book_shop.ProductServiceClient {

	return s.productService
}
