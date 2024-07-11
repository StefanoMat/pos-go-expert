package service

import (
	"context"
	"io"

	"github.com/stefanomat/pos-go-expert/14-gRPC/internal/entity"
	"github.com/stefanomat/pos-go-expert/14-gRPC/internal/infra/database"
	"github.com/stefanomat/pos-go-expert/14-gRPC/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductService struct {
	pb.UnsafeProductServiceServer
	ProductDB database.Product
}

func NewProductService(productDB database.Product) *ProductService {
	return &ProductService{
		ProductDB: productDB,
	}
}

func (p *ProductService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.Product, error) {
	product, err := entity.NewProduct(in.GetName(), in.GetPrice())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "product is invalid: %s", err.Error())
	}
	product, err = p.ProductDB.Create(product)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "an error occurred creating the product: %s", err.Error())
	}
	productResponse := &pb.Product{
		Id:    product.ID.String(),
		Name:  product.Name,
		Price: product.Price,
	}
	return productResponse, nil
}

func (p *ProductService) ListProducts(ctx context.Context, in *pb.Blank) (*pb.ProductList, error) {
	products, err := p.ProductDB.FindAll(0, 0, "")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "an error occurred listing the products: %s", err.Error())
	}
	var productsResponse []*pb.Product
	for _, product := range products {
		productResponse := &pb.Product{
			Id:    product.ID.String(),
			Name:  product.Name,
			Price: product.Price,
		}
		productsResponse = append(productsResponse, productResponse)
	}
	return &pb.ProductList{Products: productsResponse}, nil
}

func (p *ProductService) GetCategory(ctx context.Context, in *pb.ProductGetRequest) (*pb.Product, error) {
	product, err := p.ProductDB.FindByID(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "product not found: %s", err.Error())
	}
	productResponse := &pb.Product{
		Id:    product.ID.String(),
		Name:  product.Name,
		Price: product.Price,
	}
	return productResponse, nil
}

func (p *ProductService) CreateProductStream(stream pb.ProductService_CreateProductStreamServer) error {
	products := &pb.ProductList{}
	for {
		product, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(products)
		}
		if err != nil {
			return err
		}
		productEntity, err := entity.NewProduct(product.Name, product.Price)
		if err != nil {
			return err
		}
		productEntity, err = p.ProductDB.Create(productEntity)
		if err != nil {
			return err
		}
		products.Products = append(products.Products, &pb.Product{
			Id:    productEntity.ID.String(),
			Name:  productEntity.Name,
			Price: productEntity.Price,
		})
	}
}

func (p *ProductService) CreateProductStreamBidirectional(stream pb.ProductService_CreateProductStreamBidirectionalServer) error {
	for {
		product, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		productEntity, err := entity.NewProduct(product.Name, product.Price)
		if err != nil {
			return err
		}
		productEntity, err = p.ProductDB.Create(productEntity)
		if err != nil {
			return err
		}
		err = stream.Send(&pb.Product{
			Id:    productEntity.ID.String(),
			Name:  productEntity.Name,
			Price: productEntity.Price,
		})
		if err != nil {
			return err
		}

	}
}
