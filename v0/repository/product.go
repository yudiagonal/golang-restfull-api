package repository

import (
	"Go-crud-api/helper"
	"Go-crud-api/v0/model"
	"context"
	"database/sql"
)

// interface for Product repository implementation
type ProductRepository interface {
	SaveProduct(ctx context.Context, tx *sql.Tx, model model.Product) model.Product
	UpdateProduct(ctx context.Context, tx *sql.Tx, model model.Product) model.Product
	DeleteProduct(ctx context.Context, tx *sql.Tx, model model.Product)
	FindByIdProduct(ctx context.Context, tx *sql.Tx, ProductId int) (model.Product, error)
	FindAllProduct(ctx context.Context, tx *sql.Tx) []model.Product
}

// Product repository implementation
type ProductRepositoryImpl struct {
}

// new Product repository
func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

// Save implements ProductRepository.
func (repository *ProductRepositoryImpl) SaveProduct(ctx context.Context, tx *sql.Tx, product model.Product) model.Product {
	SQL := "INSERT INTO products(nama_barang,harga,jenis,meta_keyword) VALUES(?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, product.NamaBarang, product.Harga, product.Jenis, product.MetaKeyword)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.ID = uint64(id)
	return product

}

// Update implements ProductRepository.
func (*ProductRepositoryImpl) UpdateProduct(ctx context.Context, tx *sql.Tx, model model.Product) model.Product {
	panic("unimplemented")
}

// Delete implements ProductRepository.
func (*ProductRepositoryImpl) DeleteProduct(ctx context.Context, tx *sql.Tx, model model.Product) {
	panic("unimplemented")
}

// FindAll implements ProductRepository.
func (*ProductRepositoryImpl) FindAllProduct(ctx context.Context, tx *sql.Tx) []model.Product {
	panic("unimplemented")
}

// FindById implements ProductRepository.
func (*ProductRepositoryImpl) FindByIdProduct(ctx context.Context, tx *sql.Tx, ProductId int) (model.Product, error) {
	panic("unimplemented")
}
