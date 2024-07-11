package main

import (
	"net"

	"github.com/stefanomat/pos-go-expert/14-gRPC/internal/entity"
	"github.com/stefanomat/pos-go-expert/14-gRPC/internal/infra/database"
	"github.com/stefanomat/pos-go-expert/14-gRPC/internal/pb"
	"github.com/stefanomat/pos-go-expert/14-gRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)

	productService := service.NewProductService(*productDB)

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, productService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}

}
