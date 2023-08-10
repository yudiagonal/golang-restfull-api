package service

import (
	"Go-crud-api/helper"
	"Go-crud-api/log/request"
	"Go-crud-api/log/response"
	model "Go-crud-api/v0/model"
	"Go-crud-api/v0/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator"
)

type ProductService interface {
	CreateProduct(ctx context.Context, request request.ProductCreateRequest) response.ProductResponse
	UpdateProduct(ctx context.Context, request request.ProductUpdateRequest) response.ProductResponse
	DeleteProduct(ctx context.Context, ProductId int)
	FindByIdProduct(ctx context.Context, ProductId int) response.ProductResponse
	FindAllProduct(ctx context.Context) []response.ProductResponse
}

// struct for implementation Product service
type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

// CreateProduct implements ProductService.
func (service *ProductServiceImpl) CreateProduct(ctx context.Context, request request.ProductCreateRequest) response.ProductResponse {
	//validasi sebelum tx dimulai
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product := model.Product{
		NamaBarang:  request.NamaBarang,
		Harga:       request.Harga,
		Jenis:       request.Jenis,
		MetaKeyword: request.MetaKeyword,
	}
	product = service.ProductRepository.SaveProduct(ctx, tx, product)

	return helper.ToProductResponse(product)
}

// Delete implements ProductService.
func (*ProductServiceImpl) DeleteProduct(ctx context.Context, ProductId int) {
	panic("unimplemented")
}

// FindAll implements ProductService.
func (*ProductServiceImpl) FindAllProduct(ctx context.Context) []response.ProductResponse {
	panic("unimplemented")
}

// FindById implements ProductService.
func (*ProductServiceImpl) FindByIdProduct(ctx context.Context, ProductId int) response.ProductResponse {
	panic("unimplemented")
}

// Update implements ProductService.
func (*ProductServiceImpl) UpdateProduct(ctx context.Context, request request.ProductUpdateRequest) response.ProductResponse {
	panic("unimplemented")
}

// new Product service
func NewProductService(ProductRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: ProductRepository,
		DB:                DB,
		Validate:          validate,
	}
}
